package orderbook

import (
	"fmt"
	"math/rand"
)

type Order interface {
	GetType() OrderType
	GetPrice() float64
	GetVolume() int
	Fill(int) Trade
}

func NewOrder(o OrderType, price float64, v int, stock string) Order {
	return &orderImpl{
		id:       rand.Int(),
		bidOrAsk: o,
		price:    price,
		volume:   v,
		stock:    stock,
	}
}

type orderImpl struct {
	id       int
	ticker   string
	bidOrAsk OrderType // true for ask, false for bid
	price    float64
	volume   int
	stock    string
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
	return NewTrade(o.id, quantity, o.price, o.stock)
}
func (o *orderImpl) GetVolume() int {
	return o.volume
}
func (o *orderImpl) GetType() OrderType {
	return o.bidOrAsk
}
