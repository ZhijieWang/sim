package common

import (
	"fmt"
	"math/rand"
)

type Account interface {
	Commit(stock string, quantity int, price float64, BidOrAsk OrderType) (Order, error)
	Update(trade Trade)
	Cancel(id OrderId)
	GetId() AccountId
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
func (a *accountImpl) GetId() AccountId {
	return a.id
}
func (a *accountImpl) Update(trade Trade) {
	stock := trade.GetStockSymbol()
	a.cash.Update(trade)
	if _, ok := a.positions[stock]; !ok {
		a.positions[stock] = NewStockPosition(stock)
	}
	a.positions[stock].Update(trade)
}
func (a *accountImpl) Commit(stock string, quantity int, price float64, BidOrAsk OrderType) (Order, error) {
	var order Order
	var err error
	switch BidOrAsk {
	case BidOrder:
		order, err = a.cash.Commit(quantity, price, stock)
	case AskOrder:
		order, err = a.positions[stock].Commit(quantity, price, stock)
	}
	if err != nil {
		return nil, err
	} else {
		order.SetTraderId(a.id)
		return nil, nil
	}
}
