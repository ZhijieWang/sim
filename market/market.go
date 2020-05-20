package market

import (
	"marketplace/common"
	"marketplace/orderbook"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Market interface {
	//OrderHandler(Order) bool
	GetSymbol() string
	PlaceOrder(common.Order) []common.Trade
	GetQuote() common.Quote
	init(float64, int)
}

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

type marketImpl struct {
	stock      string
	shares     int
	lastPrice  float64
	lastVolume int
	book       orderbook.OrderBook
}

func NewMarket() Market {
	m := &marketImpl{
		stock:      String(10),
		shares:     1000,
		lastPrice:  0,
		lastVolume: 0,
		book:       nil,
	}
	m.init(1, 1000)
	return m
}

// init function implements the IPO process of a stock.
func (m *marketImpl) init(price float64, quantity int) {
	m.book = orderbook.NewOrderBook()
	m.book.PlaceAsk(common.NewOrder(common.BidOrder, price, quantity, m.stock))
}
func (m *marketImpl) GetSymbol() string {
	return m.stock
}
func (m *marketImpl) GetQuote() common.Quote {
	curBidP, curBidV := m.book.GetCurrentBid()
	curAskP, curAskV := m.book.GetCurrentAsk()
	return common.Quote{
		m.lastPrice,
		m.lastVolume,
		curBidP,
		curBidV,
		curAskP,
		curAskV,
	}
}

func (m *marketImpl) PlaceOrder(order common.Order) []common.Trade {
	var trades []common.Trade
	switch order.GetType() {
	case common.AskOrder:
		_, trades = m.book.PlaceAsk(order)

	case common.BidOrder:
		_, trades = m.book.PlaceBid(order)
	}
	if len(trades) != 0 {
		// trade made, need to update last price and volume
		m.lastPrice = trades[len(trades)-1].GetFilledAvgPrice()
		m.lastVolume = 0
		for _, t := range trades {
			m.lastVolume += t.GetFilledQuantity()
		}
		m.lastVolume /= 2
		// the trades come in pair (1 buy 1 sell). The volume should be executed.
	}
	return trades
}
