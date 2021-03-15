package main_test

import (
	"marketplace/common"
	"marketplace/exchange"
	"marketplace/participant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulationInitiaton(t *testing.T) {
	e := exchange.InitExchange()
	e.NewMarket(participant.NewMarketMaker("A"))
	p := participant.NewParticipant(1, 10000)
	assert.NotNil(t, p.Trade(e.GetAllQuotes()))
}

type TestTrader struct {
	account common.Account
}

func (t TestTrader) Trade(quotes map[string]common.Quote) common.Order {
	o, err := t.account.Commit("A", quotes["A"].AskVolume/10, quotes["A"].CurrentAsk, common.BidOrder)
	if err != nil {
		panic(err)
	}
	return o
}
func (t TestTrader) GetBalance() map[string]common.Balance {
	return t.account.GetBalance()
}
func TestSimulationRun(t *testing.T) {
	e := exchange.InitExchange()
	e.NewMarket(participant.NewMarketMaker("A"))
	p := TestTrader{
		common.NewDefaultAccount(100000),
	}
	o := p.Trade(e.GetAllQuotes())
	e.SubmitOrder(o)
	e.GetTrades()
	assert.Equal(t, 2, len(e.GetTrades()))
}
