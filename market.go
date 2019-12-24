package main

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type OrderStatus int

const (
	Created OrderStatus = iota
	Placed
	PartiallyFilled
	Filled
	Canceled
)

type OrderType int

const (
	Bid OrderType = iota
	Ask
)

type Quote struct {
	LastPrice  float32
	CurrentBid float32
	CurrentAsk float32
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

type Order struct {
	Timestamp time.Time
	Quantity  int
	Price     float32
	Filled    int
	Status    OrderStatus
	Origin    string
	OrderType
}

type OrderBook struct {
	Bid []Order
	Ask []Order
}
type Market interface {
	PlaceOrder(Order) bool
	GetSymbol() string
	MatchOrder() []Trade
	GetQuote() Quote
}
type MarketImpl struct {
	Stock  string
	Price  float32
	Shares int32
	Orders OrderBook
}

func NewMarket() Market {
	return &MarketImpl{
		Stock:  String(10),
		Price:  1.0,
		Shares: 1000000,
	}
}
func (m *MarketImpl) PlaceOrder(order Order) bool {
	switch order.OrderType {
	case Bid:
		m.Orders.Bid = append(m.Orders.Bid, order)
	case Ask:
		m.Orders.Ask = append(m.Orders.Ask, order)
	}
	return true
}
func (m *MarketImpl) GetQuote() Quote {
	return Quote{
		m.Price,
		0,
		0,
	}
}

type Trade struct {
	Stock        string
	Position     int // + for long, - for short
	AveragePrice float32
	Origin       string
}

func (m *MarketImpl) MatchOrder() []Trade {
	panic("Not Yet Implemented")
	return []Trade{}
}
func (m *MarketImpl) GetSymbol() string {
	return ""
}
