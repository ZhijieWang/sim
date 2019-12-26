package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// exchangeImpl implements the interface of Exchange
type exchangeImpl struct {
	Markets      map[string]Market
	Participants []Trader
}

// Exchange interface documents basic behavior of an exchange and protocol.
// Multiple echange intstances can be syndicated
type Exchange interface {
	SubmitOrder(string, Order) bool
	Receive(actor.Context)
	GetQuote(string) Quote
	GetAllQuotes() map[string]Quote
}

// InitExchange is the constructor for default exchange creation
func InitExchange() Exchange {
	e := &exchangeImpl{
		make(map[string]Market),
		[]Trader{},
	}
	for i := 1; i <= 10; i++ {
		m := NewMarket()
		e.Markets[m.GetSymbol()] = m
	}
	// for i := 1; i <= 5000000; i++ {
	// 	e.Participants = append(e.Participants, NewPariticpant())
	// }
	return e
}
func (e *exchangeImpl) executeOrder(stock string, o Order) {
	trades := e.Markets[stock].MatchOrder(o)
	actor.EmptyRootContext.Send(actor.NewLocalPID("$2"), OrderFulfillment{
		Timestamp: time.Now(),
		Trades:    trades,
		OrderID:   o.OrderID,
	})
}
func (e *exchangeImpl) SubmitOrder(Stock string, o Order) bool {
	e.Markets[Stock].PlaceOrder(o)
	go e.executeOrder(Stock, o)
	return true
}
func (e *exchangeImpl) GetQuote(stock string) Quote {
	return e.Markets[stock].GetQuote()
}

func (e *exchangeImpl) GetAllQuotes() map[string]Quote {
	retVal := map[string]Quote{}
	for k, v := range e.Markets {
		retVal[k] = v.GetQuote()
	}
	return retVal

}

func (e *exchangeImpl) NotifyExecution(T Trade) {
	context := actor.EmptyRootContext
	trader := actor.NewLocalPID(T.Origin)
	context.Request(trader, T)
}

func (e *exchangeImpl) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case SubmitOrderRequest:
		fmt.Println(context.Sender())
		success := e.SubmitOrder(msg.Stock, msg.OrderDetail)
		if success {
			fmt.Println("Order Received")
		}
		if success {
			context.Respond(OrderConfirmation{})
		}
	case GetQuoteMessage:
		if msg.Stock != "" {
			context.Respond(StockQuoteMessage{
				map[string]Quote{
					msg.Stock: e.GetQuote(msg.Stock),
				},
			})
		} else {
			resp := StockQuoteMessage{
				Value: e.GetAllQuotes(),
			}
			context.Respond(resp)
			// context.Request(context.Sender(), resp)
		}
	}
}
