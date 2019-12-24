package main

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ExchangeImpl struct {
	Markets      map[string]Market
	Participants []Trader
}
type Exchange interface {
	SubmitOrder(string, string, Order) bool
	Receive(actor.Context)
	GetQuote(string) Quote
	GetAllQuotes() map[string]Quote
}

func InitExchange() Exchange {
	e := &ExchangeImpl{
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

func (e *ExchangeImpl) SubmitOrder(T string, Stock string, o Order) bool {
	e.Markets[Stock].PlaceOrder(o)
	return true
}
func (e *ExchangeImpl) GetQuote(stock string) Quote {
	return e.Markets[stock].GetQuote()
}

func (e *ExchangeImpl) GetAllQuotes() map[string]Quote {
	retVal := map[string]Quote{}
	for k, v := range e.Markets {
		retVal[k] = v.GetQuote()
	}
	return retVal

}

func (e *ExchangeImpl) NotifyExecution(T Trade) {
	context := actor.EmptyRootContext
	trader := actor.NewLocalPID(T.Origin)
	context.Request(trader, T)
}

type SubmitOrderRequest struct {
	Stock       string
	OrderDetail Order
}
type OrderConfirmation struct {
}

func (e *ExchangeImpl) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case SubmitOrderRequest:
		success := e.SubmitOrder(context.Sender().GetAddress(), msg.Stock, msg.OrderDetail)
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
