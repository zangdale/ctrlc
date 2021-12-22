package main

import (
	"context"
	"ctrlc/db"
	"io"
	"os"
	"path/filepath"

	"net"
	"time"

	"github.com/atotto/clipboard"
	"github.com/judwhite/go-svc"
	"github.com/rs/zerolog/log"
	// "golang.design/x/clipboard"
)

func init() {
	dnsDial := net.Dialer{
		Timeout: 10 * time.Second,
	}
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dnsDial.DialContext(ctx, "udp", "223.5.5.5:53")
		},
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ser := &params{cancel: cancel, ctx: ctx}
	defer func() {
		if ser.logFile != nil {
			ser.logFile.Close()
		}
	}()
	if err := svc.Run(ser); err != nil {
		log.Panic().Err(err).Msg("svc run")
	}
}

type params struct {
	ctx     context.Context
	cancel  context.CancelFunc
	logFile *os.File
}

func (p *params) Init(env svc.Environment) error {
	log.Info().Msg("ctrlc server init")
	defer log.Info().Msg("ctrlc server init ok")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	logPath := filepath.Join(dir, "log.log")

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	p.logFile = f
	// io.MultiWriter(os.Stderr, f)
	log.Logger = log.Output(io.MultiWriter(os.Stderr, f))

	db.Init()
	return nil
}

// Start is called after Init. This method must be non-blocking.
func (p *params) Start() error {
	log.Info().Msg("ctrlc server start")
	defer log.Info().Msg("ctrlc server start ok")

	go func() {
		ticker := time.NewTicker(time.Second)
		oldText := ""
		for range ticker.C {
			select {
			case <-p.ctx.Done():
				ticker.Stop()
				return
			default:
				text, err := clipboard.ReadAll()
				if err != nil {
					log.Error().Err(err).Msg("read all")
					err = nil
				}
				// b := clipboard.Read(clipboard.FmtText)
				// text := string(b)
				if text != "" && oldText != text {
					oldText = text
					log.Debug().Msg(text)
					db.SaveCopyMsg(p.ctx, text)
				}

			}
		}
	}()

	return nil
}

// Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
// Windows Service is stopped.
func (p *params) Stop() error {
	log.Info().Msg("ctrlc server stop")
	defer log.Info().Msg("ctrlc server stop ok")
	p.cancel()
	db.Close()
	// p.logFile.Close()

	return nil
}
