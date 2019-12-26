package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/actor/middleware/opentracing"
)

// Trader type defines the interface of a basic trader.
// Actual behavior is defined in traderImpl struct
type Trader interface {
	Trade()
	Observe()
	Receive(actor.Context)
	Run()
}
type traderImpl struct {
	Self        *actor.PID
	Balance     float32
	OpenOrders  []Order
	OrderHistor []Order
	latestQuote map[string]Quote
}

func getRootContext() *actor.RootContext {
	return actor.NewRootContext(nil).WithSpawnMiddleware(opentracing.TracingMiddleware())
}
func (t *traderImpl) Run() {
	t.Observe()
	t.Trade()
}
func getExchange() *actor.PID {

	return actor.NewLocalPID("Exchange")

}
func (t *traderImpl) Observe() {
	exchange := getExchange()

	resp, err := getRootContext().RequestFuture(exchange, GetQuoteMessage{}, time.Second).Result()
	if err != nil {
		panic(err)
	}
	t.latestQuote = resp.(StockQuoteMessage).Value
}

//CreateMarketOrder should take in a Trader's ID/address in string form,
//stock symbol, quantity as int and price and order type
func (t *traderImpl) CreateMarketOrder(stock string, quantity int, price float32, bidorask OrderType) Order {
	return Order{
		Timestamp: time.Now(),
		Quantity:  quantity,
		Price:     price,
		Filled:    0,
		Status:    Created,
		OrderType: bidorask,
		Origin:    *t.Self,
		OrderID:   String(10),
	}
}
func (t *traderImpl) Trade() {
	if t.latestQuote == nil {
		panic("NO OBSERVATION")
	}
	r := rand.Intn(2)

	switch r {
	case 0:
		// return "", Order{}
		//do nothing
	case 1:
		k := getSomeKey(t.latestQuote)
		order := t.CreateMarketOrder(k, 1, t.latestQuote[k].CurrentAsk, Bid)
		result, _ := getRootContext().RequestFuture(getExchange(), SubmitOrderRequest{k, order}, time.Second).Result()
		if _, ok := result.(OrderConfirmation); !ok {
			panic("Order Submission Failed")
		} else {
			fmt.Println("Order Confirmed")
			t.OpenOrders = append(t.OpenOrders, order)
		}

	case 2:
		// if hold position, find one to sell
	}

}
func getSomeKey(m map[string]Quote) string {
	for k := range m {
		return k
	}
	return ""
}

func (t *traderImpl) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		t.Self = context.Self()

	case OrderConfirmation:
		fmt.Println(msg)
	case OrderFulfillment:
		fmt.Printf("Order Fulfillment %+v\n", msg)
	case TICK:
		t.Run()
		context.Respond(DONE{context.Self().GetAddress()})
	}
}

// NewParticipant instantiate a new trader
// TODO: instantiate NewParticipant as a SPawnFunc with autonatically keyed names
func NewPariticpant() Trader {
	t := &traderImpl{
		nil,
		10000.00,
		[]Order{},
		[]Order{},
		nil,
	}
	return t
}
