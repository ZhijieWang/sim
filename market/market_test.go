package market

import (
	"marketplace/orderbook"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketInitialization(t *testing.T) {
	m := NewMarket()
	q := m.GetQuote()
	assert.Equal(t, 0, q.LastTradedVolume)
	assert.Equal(t, 0.0, q.LastTradedPrice)
	assert.Equal(t, 0.0, q.CurrentBid)
	assert.Equal(t, 1.0, q.CurrentAsk)
}

func TestOrderExecution(t *testing.T) {

	m := NewMarket()
	trades := m.PlaceOrder(orderbook.NewOrder(orderbook.BidOrder, 1.0, 1000, ""))
	assert.Equal(t, 2, len(trades))
	assert.Equal(t, 1000, trades[0].GetFilledQuantity())
	q := m.GetQuote()
	assert.Equal(t, 1000, q.LastTradedVolume)
	assert.Equal(t, 1.0, q.LastTradedPrice)
	assert.Equal(t, math.Inf(1), q.CurrentAsk)
	assert.Equal(t, 0.0, q.CurrentBid)
}
