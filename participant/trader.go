package participant

import (
	"marketplace/common"
	"marketplace/exchange"
	"math/rand"
)

// Trader type defines the interface of a basic trader.
// Actual behavior is defined in traderImpl struct
type Trader interface {
	Trade() common.Order
	GetBalance() map[string]common.Balance
}
type traderImpl struct {
	exchange exchange.Exchange
	balance  map[string]common.Balance
	account  common.Account
}

func (t *traderImpl) Trade() common.Order {
	r := rand.Intn(3)
	quotes := t.exchange.GetAllQuotes()
	switch r {
	case 0:
		// return "", common.Order{}
		//do nothing
	case 1:
		k := getSomeKey(quotes)
		order, err := t.account.Commit(k, int(t.account.GetBalance()["$"].Available/quotes[k].CurrentAsk), quotes[k].CurrentAsk, common.AskOrder)
		if err != nil {
			panic(err)
		}

		return order

	case 2:
		// if hold position, find one to sell
		return nil
	}
	return nil
}
func (t *traderImpl) GetBalance() map[string]common.Balance {
	return t.account.GetBalance()
}
func getSomeKey(m map[string]common.Quote) string {
	for k := range m {
		return k
	}
	return ""
}

// NewParticipant instantiate a new trader
// TODO: instantiate NewParticipant as a SpawnFunc with autonatically keyed names
func NewParticipant(e exchange.Exchange) Trader {
	aid, err := e.NewAccount(1000000.0)
	if err != nil {
		panic("Not yet Implemented")
	}
	t := &traderImpl{
		e,
		e.GetBalance(aid),
		common.NewDefaultAccount(1000000.0),
	}
	return t
}

type marketMakerImpl struct {
	account      common.Account
	ticker       string
	ipo_price    float64
	ipo_quantity int
	initialized  bool
}

func NewMarketMaker(e exchange.Exchange, tickr string) Trader {
	m := &marketMakerImpl{
		account:      common.NewMarketMakerAccount(tickr, 10000),
		ticker:       tickr,
		ipo_price:    1.0,
		ipo_quantity: 10000,
		initialized:  false,
	}
	return m
}
func (m *marketMakerImpl) GetBalance() map[string]common.Balance {
	return m.account.GetBalance()
}
func (m *marketMakerImpl) Trade() common.Order {
	if m.initialized {
		return nil
	} else {

		order, err := m.account.Commit(m.ticker, m.ipo_quantity, m.ipo_price, common.AskOrder)
		if err != nil {
			panic(err)
		}
		return order
	}
}
