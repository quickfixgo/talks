package main

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/talks/whygo/app"
	"os"
	"os/signal"
)

func main() {
	app := app.NewAcceptorApp()

	cfg, _ := os.Open("whygo/cfg/acceptor.cfg")
	appSettings, _ := quickfix.ParseSettings(cfg)
	logFactory, _ := quickfix.NewFileLogFactory(appSettings)
	acceptor, _ := quickfix.NewAcceptor(app, appSettings, logFactory)
	acceptor.Start()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt)
	<-interrupt
}
