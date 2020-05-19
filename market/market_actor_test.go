package market

import (
	"testing"
)

//var rootContext *actor.RootContext = getRootContext()

//type DummyTrader struct {
//	WG     *sync.WaitGroup
//	Trades []Trade
//}

//func (d *DummyTrader) Receive(context actor.Context) {
//	switch m := context.Message().(type) {
//	case Trade:
//		d.Trades = append(d.Trades, m)
//		d.WG.Done()
//	}
//}
//
//var t1 = &DummyTrader{
//	&sync.WaitGroup{},
//	[]Trade{},
//}
//var t2 = &DummyTrader{
//	&sync.WaitGroup{},
//	[]Trade{},
//}
//
//// var a *actor.PID =
//var a, _ = rootContext.SpawnNamed(
//	actor.PropsFromProducer(func() actor.Actor {
//		return t1
//	},
//	), "D1")
//var b, _ = rootContext.SpawnNamed(
//	actor.PropsFromProducer(func() actor.Actor {
//		return t2
//	},
//	), "D2")
//var IPOOrder Order = Order{
//	Timestamp: time.Now(),
//	Quantity:  100,
//	Price:     1.20,
//	Filled:    0,
//	Status:    Placed,
//	Origin:    *a,
//	OrderType: Ask,
//	OrderID:   String(10),
//}
//
//func TestIPO(t *testing.T) {
//
//	m := NewMarket()
//	m.PlaceOrder(IPOOrder)
//	assert.Equal(t,
//		float32(1.0), m.GetQuote().LastPrice,
//	)
//	assert.Equal(
//		t,
//		float32(1.2),
//		m.GetQuote().CurrentAsk,
//	)
//	assert.Equal(
//		t,
//		float32(1.0),
//		m.GetQuote().CurrentBid,
//	)
//}
//
//func TestIPOOrderMatch(t *testing.T) {
//	// msgs = msgs[0:]
//	m := NewMarket()
//	t1.WG.Add(1)
//	t2.WG.Add(1)
//	m.PlaceOrder(IPOOrder)
//	m.PlaceOrder(Order{
//		Timestamp: time.Now(),
//		Quantity:  20,
//		Price:     1.35,
//		Filled:    0,
//		Status:    Placed,
//		Origin:    *b,
//		OrderType: Bid,
//		OrderID:   String(10),
//	})
//	t1.WG.Wait()
//	t2.WG.Wait()
//	assert.Equal(
//		t,
//		float32(1.20),
//		m.GetQuote().LastPrice,
//	)
//	assert.Equal(
//		t,
//		float32(1.35),
//		m.GetQuote().CurrentBid,
//	)
//	assert.Equal(
//		t,
//		float32(1.20),
//		m.GetQuote().CurrentAsk,
//	)
//	sum := 0
//	assert.NotEmpty(t, t2.Trades)
//	for _, i := range t2.Trades {
//		sum += i.Position
//	}
//	assert.Equal(t, 20, sum)
//}
//
func TestMarketInitialization(t *testing.T) {
	m := NewMarket()
	q = m.GetQuote()
	assertEqual(t, q, nil)
	//	m.PlaceOrder(IPOOrder)
	//	m.PlaceOrder(Order{
	//		Timestamp: time.Now(),
	//		Quantity:  20,
	//		Price:     1.35,
	//		Filled:    0,
	//		Status:    Placed,
	//		Origin:    *b,
	//		OrderType: Bid,
	//		OrderID:   String(10),
	//	})
	//	// t1.WG.Wait()
	//	// t2.WG.Wait()
	//	// assert.Equal(
	//	// 	t,
	//	// 	float32(1.20),
	//	// 	m.GetQuote().LastPrice,
	//	// )
	//	// assert.Equal(
	//	// 	t,
	//	// 	float32(1.35),
	//	// 	m.GetQuote().CurrentBid,
	//	// )
	//	// assert.Equal(
	//	// 	t,
	//	// 	float32(1.20),
	//	// 	m.GetQuote().CurrentAsk,
	//	// )
	//	// sum := 0
	//	// assert.NotEmpty(t, t2.Trades)
	//	// for _, i := range t2.Trades {
	//	// 	sum += i.Position
	//	// }
	//	// assert.Equal(t, 20, sum)
}
