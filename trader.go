package main

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
	Observe()
	Run()
}
type traderImpl struct {
	exchange  exchange.Exchange
	balance   common.Position
	accountId common.AccountId
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
		order, err := t.account.Commit(k, int(t.account.GetBalance()["$"]/quotes[k].CurrentAsk), quotes[k].CurrentAsk, common.AskOrder)
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
func getSomeKey(m map[string]common.Quote) string {
	for k := range m {
		return k
	}
	return ""
}

// NewParticipant instantiate a new trader
// TODO: instantiate NewParticipant as a SPawnFunc with autonatically keyed names
func NewParticipant(e exchange.Exchange) Trader {
	aid, err := e.NewAccount(1000000, 0)
	t := &traderImpl{
		e,
		e.GetBalance(aID),
		aid,
	}
	return t
}
