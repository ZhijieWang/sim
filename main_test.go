package main_test

import (
	"marketplace/exchange"
	"marketplace/participant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulationInitiator(t *testing.T) {
	e := exchange.InitExchange()
	e.NewMarket(participant.NewMarketMaker("A"))
	p := participant.NewParticipant(1, 10000)

	assert.NotNil(t, p.Trade(e.GetAllQuotes()))
}
