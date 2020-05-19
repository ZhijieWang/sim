package market

import (
	"marketplace/orderbook"
	"math/rand"
	"time"
)

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

type marketImpl struct {
	stock      string
	shares     int
	lastPrice  float64
	lastVolume int
	book       orderbook.OrderBook
}

func NewMarket() Market {
	return &marketImpl{}
}

// init function implements the IPO process of a stock.
func (m *marketImpl) init(price float32, quantity int) {
	m.book = orderbook.NewOrderBook(price, quantity)
}
func (m *marketImpl) GetSymbol() string {
	return m.stock
}
