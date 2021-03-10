package participant

import (
	"fmt"
	"marketplace/common"
	"marketplace/exchange"
	"math/rand"
)

// Trader type defines the interface of a basic trader.
// Actual behavior is defined in traderImpl struct
type Trader interface {
	Trade()
	GetBalance()
	GetPosition()
}
type traderImpl struct {
	exchange exchange.Exchange
	balance  map[string]common.Balance
	account  common.Account
}

func (t *traderImpl) Trade() {
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
		if t.exchange.SubmitOrder(order) {
			fmt.Println("Order Confirmed")
		} else {
			panic(fmt.Errorf("Order Failed %+v", order))
		}

	case 2:
		// if hold position, find one to sell
	}

}
func (t *traderImpl) GetBalance() float64 {
	return t.account.GetBalance()['$']
}
func getSomeKey(m map[string]common.Quote) string {
	for k := range m {
		return k
	}
	return ""
}

// NewParticipant instantiate a new trader
// TODO: instantiate NewParticipant as a SPawnFunc with autonatically keyed names
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
