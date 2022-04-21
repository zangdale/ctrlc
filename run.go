package main

import (
	"context"
	"ctrlc/db"
	"ctrlc/server"
	"ctrlc/utils"

	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	"github.com/rs/zerolog/log"
)

type params struct {
	ctx     context.Context
	cancel  context.CancelFunc
	logFile *os.File
	webUrl  string
	homeDir string
}

func (p *params) Init() error {
	log.Info().Msg("ctrlc server init")
	defer log.Info().Msg("ctrlc server init ok")

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir, err := utils.HomeDir()
	if err != nil {
		return err
	}
	p.homeDir = dir
	logPath := filepath.Join(p.homeDir, "log.log")

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	p.logFile = f
	// io.MultiWriter(os.Stderr, f)
	log.Logger = log.Output(io.MultiWriter(os.Stderr, f))

	db.Init(p.homeDir)
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
					server.AddMsg(text)
					db.SaveCopyMsg(p.ctx, text)
				}

			}
		}
	}()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	p.webUrl = fmt.Sprintf("http://%s", ln.Addr().String())
	server := &http.Server{Handler: server.Init()}
	go func() {
		e := make(chan error)
		go func(re chan error) {
			if err := server.Serve(ln); err != nil {
				re <- err
			}
		}(e)
		log.Info().Str("weburl", p.webUrl).Msg("http server")

		select {
		case <-p.ctx.Done():
			if err := server.Close(); err != nil {
				log.Error().Err(err).Msg("http close")
			}
		case err := <-e:
			log.Error().Err(err).Msg("http server")
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
