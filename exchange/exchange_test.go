package exchange

import (
	"marketplace/participant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeAccount(t *testing.T) {
	e := InitExchange()
	a := participant.NewMarketMaker("A")
	e.NewMarket(a)
	ref := e.(*exchangeImpl)
	assert.NotZero(t, len(ref.markets))
	assert.NotZero(t, len(ref.accounts))
}
