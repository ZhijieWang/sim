package market

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"math/rand"
	"time"
)

type Quote struct {
	LastTradedPrice  float32
	LastTradedVolume int
	CurrentBid       float32
	BidVolumer       int
	CurrentAsk       float32
	AskVolume        int
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

const (
	// MaxPriceRange bounds the order book price point size.
	// If an order's price is within the current price +/- max price range,
	// the order will be condiser as a valid one. Or it will be rejected.
	MaxPriceRange = 10
)

//type Entry struct {
//	Price    float32
//	Quantity int
//	Orders   []Order
//}

//type orderBook struct {
//	Entries []Entry
//	askMin  int
//	bidMax  int
//}

// Market interface declares interface for the market implementation
type Market interface {
	//OrderHandler(Order) bool
	GetSymbol() string
	//	MatchOrder(Order) []Trade
	GetQuote() Quote
	init(float32, int)
}

//marketImpl implements the basic market behavior with FIFO order execution
type marketImpl struct {
	Stock  string
	Price  float32
	Shares int32
	Orders orderBook
}

// NewMarket constructor
func NewMarket() Market {
	//	entryInitializer := func(size int, offset float32) []Entry {
	//		e := make([]Entry, size)
	//		for i := float32(0.0); int(i) < len(e); i++ {
	//			e[int(i)].Price = offset + i*0.01
	//		}
	//		return e
	//	}
	return &marketImpl{
		Stock:  String(10),
		Price:  1.0,
		Shares: 1000000,
		Orders: orderBook{
			//TODO: would need to implement growing and resizing of OrderBook due to price drift
			entryInitializer(100, 1),
			MaxPriceRange*10 - 1, // 999
			0,
		},
	}
}

//func (m *marketImpl) OrderHandler(order Order) bool {
//	switch order.OrderType {
//	case Bid:
//		m.MatchBidOrder(order)
//	case Ask:
//		m.MatchAskOrder(order)
//	}
//	return true
//}
func (m *marketImpl) GetQuote() Quote {
	return Quote{
		m.Price,
		m.Orders.Entries[m.Orders.bidMax].Price,
		m.Orders.Entries[m.Orders.askMin].Price,
	}
}

//func (m *marketImpl) placeOrder(order Order) bool {
//
//}

type TradeType int

const (
	Bought TradeType = iota
	Sold
)

// Trade defines the message of a trade to be placed
type Trade struct {
	Stock        string
	Position     int
	AveragePrice float32
	TradeType    TradeType
}

//func (o Order) Fill(q int, price float32) {
//	var t Trade
//	if o.Quantity >= q {
//		o.Quantity -= q
//		o.Filled += q
//		o.Status = Filled
//		t = Trade{
//			"stock",
//			+q,
//			price,
//			0,
//		}
//		switch o.OrderType {
//		case Bid:
//			t.TradeType = Bought
//		case Ask:
//			t.TradeType = Sold
//		}
//		NotifyExecution(t, &o.Origin)
//	}
// return t
//}

// MatchBidOrder is responsible to take a bid and match to available Asks.
// If there are available Asks, try to match the orders, up to the bid limit price limit and bid quantity.
// Startitng from the lowest ask price, match Ask order one by one.
// if the ask price is below the bid limit, match the order up to the bid quantity.
// if there are remaining bid quantity, repeat.
func (m *marketImpl) fill(o Order, e *Entry, q int) {
	o.Fill(q, e.Price)
	e.Quantity -= q
	m.Price = e.Price
	for _, i := range e.Orders {
		if q == 0 {
			break
		}
		switch {
		case q > i.Quantity:
			i.Fill(i.Quantity, e.Price)
			q -= i.Quantity
		case q <= i.Quantity:
			i.Fill(q, e.Price)
			q -= q
		}
	}
}
func (m *marketImpl) MatchBidOrder(o Order) {
	q := 0
	var askmin int
	for askmin = m.Orders.askMin; m.Orders.Entries[askmin].Price < o.Price && o.Quantity >= 0; askmin++ {
		q = m.Orders.Entries[askmin].Quantity
		switch {
		case q != 0 && q <= o.Quantity:
			m.fill(o, &m.Orders.Entries[askmin], q)
			m.Orders.askMin = askmin

		case q > o.Quantity:
			m.fill(o, &m.Orders.Entries[askmin], o.Quantity)
			m.Orders.askMin = askmin
		case q == 0:
			continue
		}
	}

	//TODO: Evaluate End condition
	// if o.Quantity == 0 {
	// 	o.MarkComplete()
	// }
	switch o.Quantity == 0 {
	case true:
		// do nothing for now
		// TODO: evaluate how to mark order filled
	case false:
		// order not matched completely, remaining order goes to the order book
		// calcuate the offset from the orderbook, insert the order.
		offset := int((o.Price - 1) * 100)
		e := m.Orders.Entries[offset]
		e.Quantity += o.Quantity
		e.Orders = append(e.Orders, o)
		if m.Orders.bidMax < offset {
			m.Orders.bidMax = offset
		}
	}
}

func (m *marketImpl) MatchAskOrder(o Order) {

	q := 0
	var bmax int
	for bmax = m.Orders.bidMax; m.Orders.Entries[bmax].Price > o.Price && q < o.Quantity; bmax-- {

		q = m.Orders.Entries[bmax].Quantity
		switch {
		case q != 0 && q <= o.Quantity:
			m.fill(o, &m.Orders.Entries[bmax], q)
			m.Orders.bidMax = bmax

		case q > o.Quantity:
			m.fill(o, &m.Orders.Entries[bmax], o.Quantity)
			m.Orders.bidMax = bmax
		case q == 0:
			continue
		}
	}
	switch o.Quantity == 0 {
	case true:
		// do nothing for now
		// TODO: evaluate how to mark order filled
	case false:
		// order not matched completely, remaining order goes to the order book
		// calcuate the offset from the orderbook, insert the order.
		offset := int((o.Price - 1) * 100)
		e := m.Orders.Entries[offset]
		e.Quantity += o.Quantity
		e.Orders = append(e.Orders, o)
		m.Orders.Entries[offset] = e
		if m.Orders.askMin > offset {
			m.Orders.askMin = offset
		}
	}
}

// init function implements the IPO process of a stock.
func (m *marketImpl) init(price float32, quantity int) {
	m.PlaceOrder(Order{
		Timestamp: time.Now(),
		Quantity:  quantity,
		Price:     price,
		Filled:    0,
		Status:    Placed,
		Origin:    *actor.NewLocalPID("MarketMaker"),
		OrderType: Ask,
		OrderID:   String(10),
	})
}

func (m *marketImpl) MatchOrder(o Order) []Trade {

	// switch o.OrderType {
	// case Bid:
	// 	return m.MatchBidOrder(o)
	// case Ask:
	// 	return m.MatchAskOrder(o)
	// }
	return []Trade{}
}
func (m *marketImpl) GetSymbol() string {
	return ""
}
