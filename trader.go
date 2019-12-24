package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/actor/middleware/opentracing"
)

type Trader interface {
	Trade()
	Observe()
	Receive(actor.Context)
	Run()
}
type TraderImpl struct {
	ID          string
	Balance     float32
	OpenOrders  []Order
	OrderHistor []Order
	latestQuote map[string]Quote
	context     *actor.RootContext
}

func (t *TraderImpl) Run() {
	t.Observe()
	t.Trade()
}
func get_exchange() *actor.PID {
	return actor.NewLocalPID("Exchange")

}
func (t *TraderImpl) Observe() {
	exchange := get_exchange()

	resp, err := t.context.RequestFuture(exchange, GetQuoteMessage{}, time.Second).Result()
	if err != nil {
		panic(err)
	}
	t.latestQuote = resp.(StockQuoteMessage).Value
}
func CreateMarketOrder(traderId string, stock string, quantity int, price float32, bidorask OrderType) Order {
	return Order{
		Timestamp: time.Now(),
		Quantity:  quantity,
		Price:     price,
		Filled:    0,
		Status:    Created,
		OrderType: bidorask,
		Origin:    traderId,
	}
}
func (t *TraderImpl) Trade() {
	if t.latestQuote == nil {
		panic("NO OBSERVATION")
	}
	r := rand.Intn(2)

	switch r {
	case 0:
		// return "", Order{}
		//do nothing
	case 1:
		k := get_some_key(t.latestQuote)
		order := CreateMarketOrder(t.ID, k, 1, t.latestQuote[k].CurrentAsk, Bid)
		result, _ := t.context.RequestFuture(get_exchange(), SubmitOrderRequest{k, order}, time.Second).Result()
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
func get_some_key(m map[string]Quote) string {
	for k := range m {
		return k
	}
	return ""
}

type GetQuoteMessage struct {
	Stock string
}
type StockQuoteMessage struct {
	Value map[string]Quote
}
type TICK struct{}
type DONE struct {
	WHO string
}

func (t *TraderImpl) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:

		t.Run()

	case OrderConfirmation:
		fmt.Println(msg)
	case TICK:
		t.Run()
		context.Respond(DONE{context.Self().GetAddress()})
	}
}

func NewPariticpant() Trader {
	t := &TraderImpl{
		String(5),
		10000.00,
		[]Order{},
		[]Order{},
		nil,
		actor.NewRootContext(nil).WithSpawnMiddleware(opentracing.TracingMiddleware()),
	}
	return t
}
