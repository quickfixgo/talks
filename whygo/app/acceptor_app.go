package app

import "github.com/quickfixgo/quickfix"
import "log"

// START BASE OMIT
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

// END BASE OMIT

// START TO FROM OMIT
func (e *AcceptorApp) ToAdmin(msgBuilder quickfix.Message,
	sessionID quickfix.SessionID) {
}

func (e *AcceptorApp) FromAdmin(msg quickfix.Message,
	sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (e *AcceptorApp) ToApp(msgBuilder quickfix.Message,
	sessionID quickfix.SessionID) error {
	return nil
}

func (e *AcceptorApp) FromApp(msg quickfix.Message,
	sessionID quickfix.SessionID) quickfix.MessageRejectError {
	log.Print("Got ", msg.String())
	return nil
}

// END TO FROM OMIT
