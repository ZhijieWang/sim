package participant

import (
	"marketplace/common"
)

// Trader type defines the interface of a basic trader.
// Actual behavior is defined in traderImpl struct
type Trader interface {
	Trade(map[string]common.Quote) common.Order
	GetBalance() map[string]common.Balance
}
type MarketMaker interface {
	Trader
	GetTicker() string
	IPO() common.Order
}
type traderImpl struct {
	account common.Account
}

func (t *traderImpl) Trade(quotes map[string]common.Quote) common.Order {
	//	r := rand.Intn(3)
	r := 1
	switch r {
	case 0:
		// return "", common.Order{}
		//do nothing
	case 1:
		k := getSomeKey(quotes)
		order, err := t.account.Commit(k, int(t.account.GetBalance()["$"].Available/quotes[k].CurrentAsk), quotes[k].CurrentAsk, common.BidOrder)
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
func NewParticipant(aid common.AccountId, balance float64) Trader {
	t := &traderImpl{
		common.NewDefaultAccount(balance),
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

func NewMarketMaker(tickr string) MarketMaker {
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
func (m *marketMakerImpl) Trade(quotes map[string]common.Quote) common.Order {
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

func (m *marketMakerImpl) IPO() common.Order {
	return m.Trade(nil)
}
func (m *marketMakerImpl) GetTicker() string {
	return m.ticker
}
