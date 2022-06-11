package main

import (
	"context"
	"ctrlc/conf"
	"ctrlc/db"
	"ctrlc/server"
	"ctrlc/utils"
	"strconv"
	"strings"

	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/atotto/clipboard"
	"github.com/rs/zerolog/log"
	"golang.design/x/clipboard"
)

type params struct {
	ctx     context.Context
	cancel  context.CancelFunc
	logFile *os.File
	webUrl  string
	webPort uint64
	homeDir string
	dbPath  string
	conf    *conf.Conf
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
	p.conf = conf.Init(p.homeDir)

	logPath := filepath.Join(p.homeDir, "log.log")

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	p.logFile = f
	// io.MultiWriter(os.Stderr, f)
	log.Logger = log.Output(io.MultiWriter(os.Stdout, f))
	p.dbPath = filepath.Join(p.homeDir, "ctrlc.db")
	db.Init(p.dbPath)
	return nil
}

// Start is called after Init. This method must be non-blocking.
func (p *params) Start() error {
	log.Info().Msg("ctrlc server start")
	defer log.Info().Msg("ctrlc server start ok")

	go func() {
		oldText := ""
		textC := clipboard.Watch(p.ctx, clipboard.FmtText)
		for b := range textC {
			// b := clipboard.Read(clipboard.FmtText)
			text := string(b)
			if text != "" && oldText != text {
				oldText = text
				log.Debug().Msg(text)
				server.AddMsg(text)
				db.SaveCopyMsg(p.ctx, text)
			}

		}
	}()

	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", p.conf.Port))
	if err != nil {
		return err
	}

	ss := strings.Split(ln.Addr().String(), ":")
	p.webPort, err = strconv.ParseUint(ss[len(ss)-1], 10, 64)
	if err != nil {
		return err
	}
	p.webUrl = fmt.Sprintf("http://127.0.0.1:%d", p.webPort)
	server := &http.Server{Handler: server.Init(p.conf.WebDir)}
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
