package common

import "math/rand"

type Trade interface {
	GetStatus() TradeStatus
	GetFilledQuantity() int
	GetFilledAvgPrice() float64
	GetOrderId() OrderId
	GetStockSymbol() string
}

type tradeImpl struct {
	tradeId int
	orderId OrderId
	filled  int
	price   float64
	status  TradeStatus
	stock   string
}

func (t tradeImpl) GetStatus() TradeStatus {
	return t.status
}

func (t tradeImpl) GetFilledQuantity() int {
	return t.filled
}
func (t tradeImpl) GetFilledAvgPrice() float64 {
	return t.price
}

func (t tradeImpl) GetOrderId() OrderId {
	return t.orderId
}
func (t tradeImpl) GetStockSymbol() string {
	return t.stock
}
func NewTrade(o Order, filled int, price float64) Trade {
	if o.GetType() == BidOrder {
		return tradeImpl{
			tradeId: rand.Int(),
			orderId: o.GetId(),
			filled:  filled,
			price:   price,
			status:  "Baught",
			stock:   o.GetSymbol(),
		}
	} else {
		return tradeImpl{
			tradeId: rand.Int(),
			orderId: o.GetId(),
			filled:  filled,
			price:   price,
			status:  "Sold",
			stock:   o.GetSymbol(),
		}
	}
}
