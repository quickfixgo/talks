package app

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42/newordersingle"
	"log"
	"strconv"
	"time"
)

//START BASE OMIT
type InitiatorApp struct {
	complete chan interface{}
}

func NewInitiatorApp() *InitiatorApp {
	return &InitiatorApp{complete: make(chan interface{})}
}

func (i *InitiatorApp) Run() {
	<-i.complete
}

//END BASE OMIT

//START OMIT
func (i *InitiatorApp) OnLogon(sessionID quickfix.SessionID) {
	log.Print("OnLogon ", sessionID)

	go func() {
		for i := 0; i < 100; i++ {
			order := newordersingle.Message{
				ClOrdID:      strconv.Itoa(i),
				HandlInst:    "1",
				Symbol:       "TSLA",
				Side:         enum.Side_BUY,
				TransactTime: time.Now(),
				OrdType:      enum.OrdType_MARKET,
			}

			log.Print("Sending Order...")
			quickfix.SendToTarget(order, sessionID)
			time.Sleep(1 * time.Second)
		}

		i.complete <- sessionID
	}()
}

//END OMIT

func (i *InitiatorApp) OnCreate(sessionID quickfix.SessionID) {
	log.Print("OnCreate ", sessionID)
}
func (i *InitiatorApp) OnLogout(sessionID quickfix.SessionID) {
	log.Print("OnLogout ", sessionID)
}
func (i *InitiatorApp) ToAdmin(msgBuilder quickfix.Message, sessionID quickfix.SessionID) {
}
func (i *InitiatorApp) ToApp(msgBuilder quickfix.Message, sessionID quickfix.SessionID) error {
	return nil
}

func (i *InitiatorApp) FromAdmin(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (i *InitiatorApp) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return
}
