package app

import "github.com/quickfixgo/quickfix"
import "log"

type AcceptorApp struct{}

func NewAcceptorApp() *AcceptorApp {
	return &AcceptorApp{}
}

func (e *AcceptorApp) OnCreate(sessionID quickfix.SessionID) {
	log.Print("OnCreate ", sessionID)
}
func (e *AcceptorApp) OnLogon(sessionID quickfix.SessionID) {
	log.Print("OnLogon ", sessionID)
}
func (e *AcceptorApp) OnLogout(sessionID quickfix.SessionID) {
	log.Print("OnLogout ", sessionID)
}
func (e *AcceptorApp) ToAdmin(msgBuilder quickfix.MessageBuilder,
	sessionID quickfix.SessionID) {
}

func (e *AcceptorApp) ToApp(msgBuilder quickfix.MessageBuilder,
	sessionID quickfix.SessionID) error {
	return nil
}

func (e *AcceptorApp) FromAdmin(msg quickfix.Message,
	sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (e *AcceptorApp) FromApp(msg quickfix.Message,
	sessionID quickfix.SessionID) quickfix.MessageRejectError {
	log.Print("Got ", msg.String())
	return nil
}
