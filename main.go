package main

import (
	"context"

	"net"
	"time"

	"github.com/getlantern/systray"
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
	if err := ser.Init(); err != nil {
		log.Panic().Err(err).Msg("ser run")
	}

	if err := ser.Start(); err != nil {
		log.Panic().Err(err).Msg("ser run")
	}

	systray.Run(onReady(ser), onExit(ser))
}
