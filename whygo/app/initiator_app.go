package app

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix42/newordersingle"
	"log"
	"time"
)

type InitiatorApp struct {
	complete chan interface{}
}

func NewInitiatorApp() *InitiatorApp {
	return &InitiatorApp{complete: make(chan interface{})}
}

func (i *InitiatorApp) Run() {
	<-i.complete
}

func (i *InitiatorApp) OnLogon(sessionID quickfix.SessionID) {
	log.Print("OnLogon ", sessionID)

	go func() {
		for i := 0; i < 100; i++ {
			order := newordersingle.Builder(
				field.NewClOrdID("100"),
				field.NewHandlInst("1"),
				field.NewSymbol("TSLA"),
				field.NewSide(enum.Side_BUY),
				&field.TransactTimeField{},
				field.NewOrdType(enum.OrdType_MARKET))

			log.Print("Sending Order...")
			quickfix.SendToTarget(order, sessionID)
			time.Sleep(1 * time.Second)
		}

		i.complete <- sessionID
	}()
}

func (i *InitiatorApp) OnCreate(sessionID quickfix.SessionID) {
	log.Print("OnCreate ", sessionID)
}
func (i *InitiatorApp) OnLogout(sessionID quickfix.SessionID) {
	log.Print("OnLogout ", sessionID)
}
func (i *InitiatorApp) ToAdmin(msgBuilder quickfix.MessageBuilder, sessionID quickfix.SessionID) {
}
func (i *InitiatorApp) ToApp(msgBuilder quickfix.MessageBuilder, sessionID quickfix.SessionID) error {
	return nil
}

func (i *InitiatorApp) FromAdmin(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (i *InitiatorApp) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return
}