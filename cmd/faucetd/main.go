package main

import (
	"database/sql"
	"flag"

	cachecash "github.com/cachecashproject/go-cachecash"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/cachecashproject/go-cachecash/faucet"
	"github.com/cachecashproject/go-cachecash/keypair"
	"github.com/cachecashproject/go-cachecash/log"
	"github.com/cachecashproject/go-cachecash/wallet"
	"github.com/cachecashproject/go-cachecash/wallet/migrations"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

var (
	configPath  = flag.String("config", "faucet.config.json", "Path to configuration file")
	keypairPath = flag.String("keypair", "faucet.keypair.json", "Path to keypair file")
	traceAPI    = flag.String("trace", "", "Jaeger API for tracing")
)

func loadConfigFile(l *logrus.Logger, path string) (*faucet.ConfigFile, error) {
	conf := faucet.ConfigFile{}
	p, err := common.NewConfigParser(l, "faucet")
	if err != nil {
		return nil, err
	}
	err = p.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf.FaucetAddr = p.GetString("faucet_addr", ":7781")
	conf.LedgerAddr = p.GetString("ledger_addr", ":7778")
	conf.Database = p.GetString("database", "faucet-wallet.db")
	conf.Insecure = p.GetInsecure()

	return &conf, nil
}

func main() {
	common.Main(mainC)
}

func connectGrpcClient(cf *faucet.ConfigFile) (ccmsg.LedgerClient, error) {
	conn, err := common.GRPCDial(cf.LedgerAddr, cf.Insecure)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial ledger service")
	}

	grpc := ccmsg.NewLedgerClient(conn)
	return grpc, nil
}

func mainC() error {
	l := log.NewCLILogger("faucetd", log.CLIOpt{JSON: true})
	flag.Parse()

	cf, err := loadConfigFile(&l.Logger, *configPath)
	if err != nil {
		return errors.Wrap(err, "failed to load configuration file")
	}

	if err := l.ConfigureLogger(); err != nil {
		return errors.Wrap(err, "failed to configure logger")
	}
	l.Info("Starting CacheCash faucetd ", cachecash.CurrentVersion)

	defer common.SetupTracing(*traceAPI, "cachecash-faucetd", &l.Logger).Flush()

	kp, err := keypair.LoadOrGenerate(&l.Logger, *keypairPath)
	if err != nil {
		return errors.Wrap(err, "failed to get keypair")
	}

	db, err := sql.Open("sqlite3", cf.Database)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}
	l.Info("opened database")

	l.Info("applying migrations")
	n, err := migrate.Exec(db, "sqlite3", migrations.Migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	l.Infof("applied %d migrations", n)

	grpc, err := connectGrpcClient(cf)
	if err != nil {
		return errors.Wrap(err, "failed to connect to ledgerd")
	}
	wallet := wallet.NewWallet(&l.Logger, kp, db, grpc)

	fs, err := faucet.NewFaucet(&l.Logger, wallet)
	if err != nil {
		return errors.Wrap(err, "failed to create faucet service")
	}

	go fs.SyncChain()

	app, err := faucet.NewApplication(&l.Logger, fs, cf)
	if err != nil {
		return errors.Wrap(err, "failed to create faucet application")
	}

	if err := common.RunStarterShutdowner(&l.Logger, app); err != nil {
		return err
	}
	return nil
}
