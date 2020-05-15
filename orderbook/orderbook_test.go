package orderbook_test

import (
	"github.com/stretchr/testify/assert"
	"marketplace/orderbook"
	"testing"
)

func TestNewOrderBook(t *testing.T) {
	book := market.NewOrderBook(1.0, 100)
	assert.NotNil(t, book)
	status := book.PlaceAsk(market.NewOrder(true, 1.0, 100))
	assert.True(t, status)
}
func TestFillTrades(t *testing.T) {
	book := market.NewOrderBook(1.0, 100)
	trades := book.Fill(market.NewOrder(true, 1.0, 100))
}
