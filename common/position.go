package common

import "fmt"

type Position interface {
	Commit(quantity int, price float64, stock string) (Order, error)
	Update(Trade) bool
	GetAvailable() float64
}

type stockPositionImpl struct {
	stock     string
	cleared   int
	committed int
}

func NewStockPosition(stock) {
	return &stockPositionImpl{
		stock:     stock,
		cleared:   0,
		committed: 0,
	}
}
func (p *stockPositionImpl) Commit(quantity int, price float64, stock string) (Order, error) {
	if stock != p.stock {
		return nil, fmt.Errorf("Commit order to wrong position. Current Position %s, Commiting to %s", p.stock, stock)
	}
	if float64(quantity) > p.GetAvailable() {
		return nil, fmt.Errorf("Does not have enough shares to cover the trade. Available %f, Commiting %d", p.GetAvailable(), quantity)
	} else {
		p.committed += quantity
		return NewOrder(AskOrder, price, quantity, p.stock), nil
	}
}
func (p *stockPositionImpl) GetAvailable() float64 {
	return float64(p.cleared - p.committed)
}

func (p *stockPositionImpl) Update(t Trade) bool {
	switch t.GetStatus() {
	case "Baught":
		p.cleared += t.GetFilledQuantity()
		p.commited -= t.GetFilledQuantity()
		return true
	case "Sold":
		p.cleared -= t.GetFilledQuantity()
		p.committed -= t.GetFilledQuantity()
		return true
	default:
		panic(fmt.Errorf("Not Yet Implemented for Trade Type %s", t.GetStatus()))
	}
}

type cashPositionImpl struct {
	cleared   float64
	committed float64
}

func NewCashPosition(money float64) Position {
	return &cashPositionImpl{
		cleared:   money,
		committed: 0.0,
	}

}
func (cash *cashPositionImpl) Commit(quantity int, price float64, stock string) (Order, error) {
	value := float64(quantity) * price
	if value > cash.GetAvailable() {
		return nil, fmt.Errorf("Cash Balance Will Reach Negative. Available %f, Commmiting %f", cash.GetAvailable(), float64(quantity)*price)
	} else {
		cash.committed += value
		return NewOrder(BidOrder, price, quantity, stock), nil
	}
}
func (cash *cashPositionImpl) GetAvailable() float64 {
	return cash.cleared - cash.committed
}

func (cash *cashPositionImpl) Update(t Trade) bool {
	value := float64(t.GetFilledQuantity()) * t.GetFilledAvgPrice()
	switch t.GetStatus() {

	case "Baught":
		cash.cleared -= value
		cash.committed -= value
		return true
	case "Sold":
		cash.cleared += value
		cash.committed -= value
		return true
	default:
		panic(fmt.Errorf("Not Yet Implemented for Trade Type %s", t.GetStatus()))

	}
}
