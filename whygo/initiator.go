package main

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/talks/whygo/app"
	"os"
)

func main() {
	app := app.NewInitiatorApp()

	cfg, _ := os.Open("whygo/cfg/initiator.cfg")
	appSettings, _ := quickfix.ParseSettings(cfg)
	logFactory, _ := quickfix.NewFileLogFactory(appSettings)
	initiator, _ := quickfix.NewInitiator(app, appSettings, logFactory)

	initiator.Start()
	app.Run()
}
