package main

import (
	"context"
	"net"
	"net/http"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/log"

	"github.com/starudream/mirai-chatgpt-bot/route"
)

func init() {
	config.SetDefault("addr", ":80")
}

func main() {
	app.Init(initServer)
	app.Add(serve)
	app.Defer(shutdown)
	err := app.Go()
	if err != nil {
		log.Fatal().Msgf("app init fail: %v", err)
	}
}

var server *http.Server

func initServer() error {
	server = &http.Server{Addr: config.GetString("addr"), Handler: route.Handler()}
	return nil
}

func serve(context.Context) error {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}
	log.Info().Msgf("http server start at %s", server.Addr)
	return server.Serve(ln)
}

func shutdown() {
	_ = server.Shutdown(context.Background())
	log.Info().Msgf("http server shutdown")
}
