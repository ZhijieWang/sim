package main

import (
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

const (
	// MaxPriceRange bounds the order book price point size.
	// If an order's price is within the current price +/- max price range,
	// the order will be condiser as a valid one. Or it will be rejected.
	MaxPriceRange = 10
)

type Entry struct {
	Price    float32
	Quantity int
}

type orderBook struct {
	Bids   []Entry
	Asks   []Entry
	askMin int
	bidMax int
}

// Market interface declares interface for the market implementation
type Market interface {
	PlaceOrder(Order) bool
	GetSymbol() string
	MatchOrder(Order) []Trade
	GetQuote() Quote
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
	return &marketImpl{
		Stock:  String(10),
		Price:  1.0,
		Shares: 1000000,
		Orders: orderBook{
			make([]Entry, MaxPriceRange*100),
			make([]Entry, MaxPriceRange*100),
			0,
			0,
		},
	}
}
func (m *marketImpl) fillAtPrice(PP int, bora OrderType) []Trade {
	panic("Not Yet Implemented")
}
func (m *marketImpl) fillPartialAtPrice(PP int, Q int, bora OrderType) []Trade {
	panic("Not Yet Implemented")
}
func (m *marketImpl) PlaceOrder(order Order) bool {
	// switch order.OrderType {
	// case Bid:
	// 	m.Orders.Bids <- order
	// case Ask:
	// 	m.Orders.Asks <- order
	// }
	return true
}
func (m *marketImpl) GetQuote() Quote {
	return Quote{
		m.Price,
		0,
		0,
	}
}

// Trade defines the message of a trade to be placed
type Trade struct {
	Stock        string
	Position     int // + for long, - for short
	AveragePrice float32
	Origin       string
}

func (m *marketImpl) MatchBidOrder(o Order) []Trade {

	totalQuant := 0
	trades := []Trade{}
	for amin := m.Orders.askMin; m.Orders.Asks[amin].Price < o.Price; amin++ {
		q := totalQuant + m.Orders.Asks[amin].Quantity
		switch {

		case q == o.Quantity:
			trades = append(trades, m.fillAtPrice(amin, Bid)...)
			break
		case q < o.Quantity:
			trades = append(trades, m.fillAtPrice(amin, Bid)...)
		case q > o.Quantity:
			trades = append(trades, m.fillPartialAtPrice(amin, o.Quantity-totalQuant, Bid)...)
		}

	}
	return trades
}

func (m *marketImpl) MatchAskOrder(o Order) []Trade {
	totalQuant := 0
	trades := []Trade{}
	for bmax := m.Orders.bidMax; m.Orders.Bids[bmax].Price > o.Price; bmax-- {
		q := totalQuant + m.Orders.Bids[bmax].Quantity
		switch {

		case q == o.Quantity:
			trades = append(trades, m.fillAtPrice(bmax, Ask)...)
			break
		case q < o.Quantity:
			trades = append(trades, m.fillAtPrice(bmax, Ask)...)
		case q > o.Quantity:
			trades = append(trades, m.fillPartialAtPrice(bmax, o.Quantity-totalQuant, Ask)...)
		}

	}
	return trades
}

func (m *marketImpl) MatchOrder(o Order) []Trade {

	switch o.OrderType {
	case Bid:
		return m.MatchBidOrder(o)
	case Ask:
		return m.MatchAskOrder(o)
	}
	return []Trade{}
}
func (m *marketImpl) GetSymbol() string {
	return ""
}
