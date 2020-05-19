package orderbook_test

import (
	"github.com/stretchr/testify/assert"
	"marketplace/orderbook"
	"testing"
)

func TestOrderExecution(t *testing.T) {
	book := orderbook.NewOrderBook(1.0, 100)
	status, trades := book.PlaceBid(orderbook.NewOrder(orderbook.BidOrder, 1.0, 100))
	price, unit := book.GetCurrentBid()
	assert.Equal(t, 0.0, price)
	assert.Equal(t, 0, unit)
	price, unit = book.GetCurrentAsk()
	assert.Equal(t, 0.0, price)
	assert.Equal(t, 0, unit)
	assert.True(t, status)
	assert.Equal(t, 100, trades[0].GetFilledQuantity())
	assert.Equal(t, 2, len(trades))
	status, trades = book.PlaceBid(orderbook.NewOrder(orderbook.BidOrder, 1.0, 1000))
	assert.Equal(t, true, status)
	assert.Nil(t, trades)

	price, unit = book.GetCurrentBid()
	assert.Equal(t, 1.0, price)
	assert.Equal(t, 1000, unit)

}

func TestOrderbookPlaceOrder(t *testing.T) {
	book := orderbook.NewOrderBook(1.0, 100)

	status, _ := book.PlaceAsk(orderbook.NewOrder(orderbook.AskOrder, 1.0, 100))
	assert.True(t, status)
	price, unit := book.GetCurrentAsk()
	assert.Equal(t, 1.0, price)
	assert.Equal(t, 200, unit)
	status, _ = book.PlaceBid(orderbook.NewOrder(orderbook.BidOrder, 0.9, 100))
	price, unit = book.GetCurrentBid()
	assert.Equal(t, 0.9, price)
	assert.Equal(t, 100, unit)
}
func TestNewOrderbook(t *testing.T) {
	book := orderbook.NewOrderBook(1.0, 100)
	assert.NotNil(t, book)
}
