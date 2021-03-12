package market

import (
	"marketplace/common"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketInitialization(t *testing.T) {
	m := NewMarket("A")
	q := m.GetQuote()
	assert.Equal(t, 0, q.LastTradedVolume)
	assert.Equal(t, 0.0, q.LastTradedPrice)
	assert.Equal(t, 0.0, q.CurrentBid)
	assert.Equal(t, math.Inf(1), q.CurrentAsk)
}
func TestMarketPlaceOrder(t *testing.T) {
	m := NewMarket("A")
	m.PlaceOrder(common.NewOrder(common.AskOrder, 1.0, 1000, "A"))
	q := m.GetQuote()
	assert.Equal(t, 0, q.LastTradedVolume)
	assert.Equal(t, 0.0, q.LastTradedPrice)
	assert.Equal(t, 0.0, q.CurrentBid)
	assert.Equal(t, 1.0, q.CurrentAsk)
}

func TestOrderExecution(t *testing.T) {

	m := NewMarket("A")
	m.PlaceOrder(common.NewOrder(common.AskOrder, 1.0, 1000, "A"))
	trades := m.PlaceOrder(common.NewOrder(common.BidOrder, 1.0, 1000, "A"))
	assert.Equal(t, 2, len(trades))
	assert.Equal(t, 1000, trades[0].GetFilledQuantity())
	q := m.GetQuote()
	assert.Equal(t, 1000, q.LastTradedVolume)
	assert.Equal(t, 1.0, q.LastTradedPrice)
	assert.Equal(t, math.Inf(1), q.CurrentAsk)
	assert.Equal(t, 0.0, q.CurrentBid)
}
