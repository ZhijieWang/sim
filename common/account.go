package common

import (
	"fmt"
	"math/rand"
)

type Account interface {
	Commit(stock string, quantity int, price float64, BidOrAsk OrderType) (Order, error)
	Update(trade Trade)
	Cancel(id OrderId)
}

type AccountId int
type accountImpl struct {
	id        AccountId
	orders    []Order
	positions map[string]Position
	cash      Position
}

func NewDefaultAccount(cash float64) Account {

	return &accountImpl{
		id:        AccountId(rand.Int()),
		orders:    []Order{},
		positions: make(map[string]Position),
		cash:      NewCashPosition(cash),
	}
}
func (a *accountImpl) Cancel(id OrderId) {
	panic(fmt.Errorf("Not yet Implemented"))
}
func (a *accountImpl) Update(trade Trade) {
	stock := trade.GetStockSymbol()
	a.cash.Update(trade)
	if val, ok := a.positions[stock]; !ok {
		a.positions[stock] = NewStockPosition()
	}
	a.positions[stock].Update(trade)
}
func (a *accountImpl) Commit(stock string, quantity int, price float64, BidOrAsk OrderType) (Order, error) {
	switch BidOrAsk {
	case BidOrder:
		a.cash.Commit(quantity, price, stock)

	case AskOrder:

		return a.positions[stock].Commit(quantity, price, stock)
	}
	return nil, nil
}
