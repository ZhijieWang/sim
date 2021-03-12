package main_test

import (
	"fmt"
	"marketplace/exchange"
	"marketplace/participant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulationInitiator(t *testing.T) {
	e := exchange.InitExchange()
	p := participant.NewParticipant(e)
	fmt.Printf("%+v\n", e)
	p.Trade()
	fmt.Printf("%+v\n", p.GetBalance())
	assert.Nil(t, nil)
}
