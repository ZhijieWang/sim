package exchange

import (
	"fmt"
	"marketplace/common"
	"marketplace/market"
	"marketplace/orderbook"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// exchangeImpl implements the interface of Exchange
type exchangeImpl struct {
	Markets      map[string]market.Market
	Participants []orderbook.Trade
}

// Exchange interface documents basic behavior of an exchange and protocol.
// Multiple echange intstances can be syndicated
type Exchange interface {
	SubmitOrder(common.Order) bool
	Receive(actor.Context)
	GetQuote(string) market.Quote
	GetAllQuotes() map[string]market.Quote
}

// InitExchange is the constructor for default exchange creation
func InitExchange() Exchange {
	e := &exchangeImpl{
		make(map[string]Market),
		[]Trader{},
	}
	for i := 1; i <= 1; i++ {
		m := NewMarket()
		e.Markets[m.GetSymbol()] = m
	}
	return e
}
func (e *exchangeImpl) executeOrder(stock string, o Order) {
	trades := e.Markets[stock].MatchOrder(o)
	actor.EmptyRootContext.Send(&o.Origin, OrderFulfillment{
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

func NotifyExecution(T Trade, trader *actor.PID) {

	context := getRootContext()
	// fmt.Printf("Notifying Execution %+v to Trader %s", T, trader.GetId())
	context.Request(trader, T)
}

func (e *exchangeImpl) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case SubmitOrderRequest:
		success := e.SubmitOrder(msg.Stock, msg.OrderDetail)
		if success {
			fmt.Printf("Order Received %+v", msg.OrderDetail)
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
