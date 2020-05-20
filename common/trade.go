package common

import "math/rand"

type Trade interface {
	GetStatus() TradeStatus
	GetFilledQuantity() int
	GetFilledAvgPrice() float64
	GetOrderId() int
	GetStockSymbol() string
}

type tradeImpl struct {
	tradeId int
	orderId int
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

func (t tradeImpl) GetOrderId() int {
	return t.orderId
}
func (t tradeImpl) GetStockSymbol() string {
	return t.stock
}
func NewTrade(oId int, filled int, price float64, stock string) Trade {
	return tradeImpl{
		tradeId: rand.Int(),
		orderId: oId,
		filled:  filled,
		price:   price,
		status:  "Filled",
		stock:   stock,
	}

}
