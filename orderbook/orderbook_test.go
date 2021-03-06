package orderbook_test

import (
	"marketplace/common"
	"marketplace/orderbook"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderExecution(t *testing.T) {
	book := orderbook.NewOrderBook()
	status, trades := book.PlaceAsk(common.NewOrder(common.AskOrder, 1.0, 100, ""))

	status, trades = book.PlaceBid(common.NewOrder(common.BidOrder, 1.0, 100, ""))
	price, unit := book.GetCurrentBid()
	assert.Equal(t, 0.0, price)
	assert.Equal(t, 0, unit)
	price, unit = book.GetCurrentAsk()
	assert.Equal(t, math.Inf(1), price)
	assert.Equal(t, 0, unit)
	assert.True(t, status)
	assert.Equal(t, 100, trades[0].GetFilledQuantity())
	assert.Equal(t, 2, len(trades))
	status, trades = book.PlaceBid(common.NewOrder(common.BidOrder, 1.0, 1000, ""))
	assert.Equal(t, true, status)
	assert.Nil(t, trades)

	price, unit = book.GetCurrentBid()
	assert.Equal(t, 1.0, price)
	assert.Equal(t, 1000, unit)

}

func TestOrderbookPlaceOrder(t *testing.T) {
	book := orderbook.NewOrderBook()

	status, _ := book.PlaceAsk(common.NewOrder(common.AskOrder, 1.0, 100, ""))
	assert.True(t, status)
	price, unit := book.GetCurrentAsk()
	assert.Equal(t, 1.0, price)
	assert.Equal(t, 100, unit)
	status, _ = book.PlaceBid(common.NewOrder(common.BidOrder, 0.9, 100, ""))
	price, unit = book.GetCurrentBid()
	assert.Equal(t, 0.9, price)
	assert.Equal(t, 100, unit)
}
func TestNewOrderbook(t *testing.T) {
	book := orderbook.NewOrderBook()
	assert.NotNil(t, book)
}
