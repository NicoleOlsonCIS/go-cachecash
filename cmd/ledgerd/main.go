package main

import (
	"context"
	"database/sql"
	"flag"
	_ "net/http/pprof"
	"time"

	cachecash "github.com/cachecashproject/go-cachecash"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/cachecashproject/go-cachecash/dbtx"
	"github.com/cachecashproject/go-cachecash/keypair"
	"github.com/cachecashproject/go-cachecash/ledgerservice"
	"github.com/cachecashproject/go-cachecash/ledgerservice/migrations"
	"github.com/cachecashproject/go-cachecash/log"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

var (
	configPath   = flag.String("config", "ledger.config.json", "Path to configuration file")
	keypairPath  = flag.String("keypair", "ledger.keypair.json", "Path to keypair file")
	traceAPI     = flag.String("trace", "", "Jaeger API for tracing")
	mineBlocks   = flag.Bool("mine-blocks", false, "Create new blocks at an interval")
	mineInterval = flag.Int("mine-interval", 300, "Create a new block ever X seconds")
)

func loadConfigFile(l *logrus.Logger, path string) (*ledgerservice.ConfigFile, error) {
	conf := ledgerservice.ConfigFile{}
	p, err := common.NewConfigParser(l, "ledger")
	if err != nil {
		return nil, err
	}
	err = p.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf.LedgerProtocolAddr = p.GetString("ledger_addr", ":7778")
	conf.StatusAddr = p.GetString("status_addr", ":8100")
	conf.Database = p.GetString("database", "host=ledger-db port=5432 user=postgres dbname=ledger sslmode=disable")
	conf.Insecure = p.GetInsecure()

	return &conf, nil
}

func main() {
	common.Main(mainC)
}

func mainC() error {
	l := log.NewCLILogger("ledgerd", log.CLIOpt{JSON: true})
	flag.Parse()

	cf, err := loadConfigFile(&l.Logger, *configPath)
	if err != nil {
		return errors.Wrap(err, "failed to load configuration file")
	}

	if err := l.ConfigureLogger(); err != nil {
		return errors.Wrap(err, "failed to configure logger")
	}
	l.Info("Starting CacheCash ledgerd ", cachecash.CurrentVersion)

	defer common.SetupTracing(*traceAPI, "cachecash-ledgerd", &l.Logger).Flush()

	kp, err := keypair.LoadOrGenerate(&l.Logger, *keypairPath)
	if err != nil {
		return errors.Wrap(err, "failed to get keypair")
	}

	db, err := sql.Open("postgres", cf.Database)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}

	// Connect to the database.
	deadline := time.Now().Add(5 * time.Minute)
	for {
		err = db.Ping()

		if err == nil {
			// connected successfully
			break
		} else if time.Now().Before(deadline) {
			// connection failed, try again
			l.Info("Connection failed, trying again shortly")
			time.Sleep(250 * time.Millisecond)
		} else {
			// connection failed too many times, giving up
			return errors.Wrap(err, "database ping failed")
		}
	}
	l.Info("connected to database")

	l.Info("applying migrations")
	n, err := migrate.Exec(db, "postgres", migrations.Migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	l.Infof("applied %d migrations", n)

	var newTxChan *(chan struct{})
	if *mineBlocks {
		ctx := dbtx.ContextWithExecutor(context.Background(), db)

		storage := ledgerservice.NewLedgerDatabase()
		lm, err := ledgerservice.NewLedgerMiner(&l.Logger, storage, kp)
		if err != nil {
			return errors.Wrap(err, "failed to create ledger miner")
		}
		lm.Interval = time.Duration(*mineInterval)
		newTxChan = &lm.NewTxChan

		err = lm.SetupCurrentBlock(ctx, 420000000)
		if err != nil {
			return errors.Wrap(err, "failed to setup current block for mining")
		}
		go lm.Run(ctx)
	}

	ls, err := ledgerservice.NewLedgerService(&l.Logger, kp, newTxChan)
	if err != nil {
		return errors.Wrap(err, "failed to create ledger service")
	}

	app, err := ledgerservice.NewApplication(&l.Logger, ls, db, cf)
	if err != nil {
		return errors.Wrap(err, "failed to create cache application")
	}

	if err := common.RunStarterShutdowner(&l.Logger, app); err != nil {
		return err
	}
	return nil
}
