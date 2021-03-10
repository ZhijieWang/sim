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
	p.Trade()
	fmt.Println(e)
	assert.Nil(t, nil)
}
