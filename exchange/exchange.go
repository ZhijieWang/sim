package exchange

import (
	"fmt"
	"marketplace/common"
	"marketplace/market"
	"marketplace/participant"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Exchange interface {
	SubmitOrder(common.Order) bool
	GetQuote(string) common.Quote
	GetAllQuotes() map[string]common.Quote
	NewAccount(float64) (common.AccountId, error)
	GetBalance(common.AccountId) map[string]common.Balance
	GetTrades() []common.Trade
	NewMarket(participant.MarketMaker)
}

type ExchangeActor struct {
	e Exchange
}

func (e *ExchangeActor) Receive(context actor.Context) {
	panic("Not yet implemented")
}

func InitExchangeActor() *ExchangeActor {
	return &ExchangeActor{
		e: InitExchange(),
	}
}

// InitExchange is the constructor for default exchange creation
func InitExchange() Exchange {
	e := &exchangeImpl{
		make(map[string]market.Market),
		make(map[common.AccountId]common.Account),
		[]common.Trade{},
	}
	return e
}

func (e *exchangeImpl) GetTrades() []common.Trade {
	return e.trades
}

type exchangeImpl struct {
	markets  map[string]market.Market
	accounts map[common.AccountId]common.Account
	trades   []common.Trade
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

func (e *exchangeImpl) NewAccount(money float64) (common.AccountId, error) {
	if money <= 0 {
		return 0, fmt.Errorf("Invalid Account Registration with %f money", money)
	}
	a := common.NewDefaultAccount(money)
	e.accounts[a.GetId()] = a
	return a.GetId(), nil
}
func (e *exchangeImpl) GetBalance(id common.AccountId) map[string]common.Balance {
	return e.accounts[id].GetBalance()
}

func (e *exchangeImpl) SubmitOrder(order common.Order) bool {
	// exchange should verify if the bid order is valid given current account
	// positions.
	// each market is responsiblle for validating ask orders.
	var trades []common.Trade
	if val, ok := e.markets[order.GetSymbol()]; ok {
		trades = val.PlaceOrder(order)

		for _, t := range trades {
			act := t.GetOrderId().AccountId
			e.accounts[act].Update(t)
		}
		e.trades = append(e.trades, trades...)
		return true
	} else {
		panic("Not Yet Implemented")
		//return false
	}
}
func (e *exchangeImpl) NewMarket(maker participant.MarketMaker) {
	m := market.NewMarket(maker.GetTicker())
	e.markets[maker.GetTicker()] = m
}
