package exchange

import (
	"marketplace/common"
	"marketplace/market"
)

type Exchange interface {
	SubmitOrder(common.Order) bool
	GetQuote(string) common.Quote
	GetAllQuotes() map[string]common.Quote
}

// InitExchange is the constructor for default exchange creation
func InitExchange() Exchange {
	e := &exchangeImpl{
		make(map[string]market.Market),
	}
	for i := 1; i <= 1; i++ {
		m := market.NewMarket()
		e.markets[m.GetSymbol()] = m
	}
	return e
}

type exchangeImpl struct {
	markets map[string]market.Market
}

func (e *exchangeImpl) GetQuote(stock string) common.Quote {
	return e.markets[stock].GetQuote()
}

func (e *exchangeImpl) GetAllQuotes() map[string]common.Quote {
	retVal := map[string]common.Quote{}
	for k, v := range e.markets {
		retVal[k] = v.GetQuote()
	}
	return retVal

}
func (e *exchangeImpl) SubmitOrder(order common.Order) bool {
	if val, ok := e.markets[order.GetSymbol()]; ok {
		val.PlaceOrder(order)
		return true
	} else {
		panic("Not Yet Implemented")
		return false
	}
}
