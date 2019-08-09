package log

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/cachecashproject/go-cachecash/common"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Heartbeat -- if set to true, will attempt to deliver the logs every
// TickInterval. Set this to false to not deliver logs at all, and instead just
// write them. Useful for testing.
var Heartbeat = true

// TickInterval is a singleton for how long to wait for a log file to fill
// before delivering it.
var TickInterval = time.Second

// Client is a logging client that uses grpc to send a structured log.
type Client struct {
	service       string
	logDir        string
	logFile       *os.File
	logLock       sync.Mutex
	ticker        *time.Ticker
	logPipeClient LogPipeClient

	heartbeatCancel context.CancelFunc

	errorMutex sync.RWMutex
	Error      error
}

// NewClient creates a new client.
func NewClient(serverAddress, service, logDir string) (*Client, error) {
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return nil, err
	}

	c := &Client{
		service: service,
		logDir:  logDir,
	}

	if Heartbeat {
		conn, err := common.GRPCDial(serverAddress)
		if err != nil {
			return nil, err
		}

		c.logPipeClient = NewLogPipeClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		c.heartbeatCancel = cancel

		c.ticker = time.NewTicker(TickInterval)

		go c.heartbeat(ctx)
	}

	return c, c.makeLog(true)
}

// Close closes any logfile and connections
func (c *Client) Close() error {
	if c.heartbeatCancel != nil {
		c.heartbeatCancel()
	}

	c.logLock.Lock()
	defer c.logLock.Unlock()

	if c.logFile != nil {
		lf := c.logFile
		c.logFile = nil
		if err := lf.Sync(); err != nil {
			return err
		}

		if err := lf.Close(); err != nil {
			return err
		}
	}

	if err := c.makeLog(false); err != nil {
		return err
	}

	c.errorMutex.RLock()
	defer c.errorMutex.RUnlock()
	return c.Error
}

func (c *Client) heartbeat(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-c.ticker.C:
			if err := c.makeLog(true); err != nil {
				c.errorMutex.Lock()
				defer c.errorMutex.Unlock()
				c.Error = errors.Errorf("Cannot make new log; canceling heartbeat. Please create a new client. Error: %v", err)
				return
			}

			if err := c.deliverLog(ctx); err != nil {
				logrus.Errorf("Received error while delivering log: %v\n", err)
				continue
			}
		}
	}
}

func (c *Client) deliverLog(ctx context.Context) error {
	return filepath.Walk(c.logDir, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p == c.logDir {
			return nil
		}

		if fi.IsDir() {
			// we do not descend.
			logrus.Debugf("Directory (%v) found in log dir, skipping\n", p)
			return filepath.SkipDir
		}

		if fi.Mode()&os.ModeType != 0 {
			// irregular file, eject
			logrus.Debugf("Irregular file (%v) found in log dir, skipping\n", p)
			return nil
		}

		c.logLock.Lock()
		if c.logFile != nil {
			if path.Base(p) == path.Base(c.logFile.Name()) {
				c.logLock.Unlock()
				// we don't want to operate on the open file
				return nil
			}
			c.logLock.Unlock()
		}

		if err := c.sendLog(ctx, p); err != nil {
			logrus.Errorf("Could not deliver log bundle %v; will retry at next heartbeat. Error: %v", p, err)
			return nil
		}

		return nil
	})
}

func (c *Client) sendLog(ctx context.Context, p string) (retErr error) {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()
	defer func() {
		// if we cannot send or otherwise operate, do not remove the file, just log
		// the error and return; we will retry on the next heartbeat.
		if retErr != nil {
			logrus.Error(retErr)
			return
		}

		// remove the file if we received no error so it can't be re-delivered.
		if err := os.Remove(p); err != nil {
			logrus.Error(err)
		}
	}()

	client, err := c.logPipeClient.ReceiveLogs(ctx)
	if err != nil {
		return err
	}

	buf := make([]byte, 2*1024*1024)

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				if _, err := client.CloseAndRecv(); err != nil && err != io.EOF {
					return err
				}

				return nil
			}
		}

		if err := client.Send(&LogData{Data: buf[:n]}); err != nil {
			return err
		}
	}
}

func (c *Client) makeLog(takeLock bool) error {
	if takeLock {
		c.logLock.Lock()
		defer c.logLock.Unlock()
	}

	f, err := ioutil.TempFile(c.logDir, "")
	if err != nil {
		return err
	}

	if c.logFile != nil {
		c.logFile.Close()
	}

	c.logFile = f
	return nil
}

func (c *Client) Write(e *logrus.Entry) error {
	c.errorMutex.RLock()
	if c.Error != nil {
		defer c.errorMutex.RUnlock()
		return c.Error
	}
	c.errorMutex.RUnlock()

	f := &types.Struct{Fields: map[string]*types.Value{}}
	for key, value := range e.Data {
		var v string
		switch value.(type) {
		case string:
			v = value.(string)
		default:
			v = fmt.Sprintf("%v", value)
		}

		f.Fields[key] = &types.Value{Kind: &types.Value_StringValue{StringValue: v}}
	}

	t, err := types.TimestampProto(e.Time)
	if err != nil {
		return err
	}

	eOut := &Entry{
		Level:   int64(e.Level),
		Fields:  f,
		Message: e.Message,
		At:      t,
		Service: c.service,
	}

	buf, err := eOut.Marshal()
	if err != nil {
		return err
	}

	c.logLock.Lock()
	defer c.logLock.Unlock()

	if c.logFile == nil {
		if err := c.makeLog(true); err != nil {
			return err
		}
	}

	if err := binary.Write(c.logFile, binary.BigEndian, int64(len(buf))); err != nil {
		return err
	}

	if _, err := c.logFile.Write(buf); err != nil {
		return err
	}

	return nil
}
