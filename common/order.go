package common

import (
	"fmt"
	"math/rand"
	"time"
)

type TradeStatus string
type OrderType int

const (
	BidOrder OrderType = iota
	AskOrder
)

type Order interface {
	GetType() OrderType
	GetPrice() float64
	GetVolume() int
	Fill(int) Trade
	GetId() OrderId
	GetSymbol() string
	SetTraderId(AccountId) bool
}

func NewOrder(o OrderType, price float64, v int, stock string) Order {
	return &orderImpl{
		id:       OrderId{0, rand.Int()},
		bidOrAsk: o,
		price:    price,
		volume:   v,
		stock:    stock,
		tid:      0,
	}
}

type OrderId struct {
	AccountId AccountId
	OrderId   int
}
type orderImpl struct {
	id        OrderId
	Timestamp time.Time
	ticker    string
	bidOrAsk  OrderType // true for ask, false for bid
	price     float64
	volume    int
	stock     string
	tid       AccountId
}

func (o *orderImpl) GetTraderId() AccountId {
	return o.tid
}
func (o *orderImpl) SetTraderId(id AccountId) bool {
	if o.tid != 0 {
		o.tid = id
		return true
	}
	return false
}
func (o *orderImpl) GetSymbol() string {
	return o.stock
}
func (o *orderImpl) GetId() OrderId {
	return o.id
}
func (o *orderImpl) GetPrice() float64 {
	return o.price
}
func (o *orderImpl) Fill(quantity int) Trade {
	if quantity > o.volume {
		panic(fmt.Errorf("Filling more than what the order has. Want %d, Has %d", quantity, o.volume))
	}
	o.volume -= quantity
	// removal of filled order is at entry level
	return NewTrade(o, quantity, o.price)
}
func (o *orderImpl) GetVolume() int {
	return o.volume
}
func (o *orderImpl) GetType() OrderType {
	return o.bidOrAsk
}
