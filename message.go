package main

import "time"

import "github.com/AsynkronIT/protoactor-go/actor"

type SubmitOrderRequest struct {
	Stock       string
	OrderDetail Order
}
type OrderConfirmation struct {
}
type OrderFulfillment struct {
	Timestamp time.Time
	Trades    []Trade
	OrderID   string
}
type GetQuoteMessage struct {
	Stock string
}
type StockQuoteMessage struct {
	Value map[string]Quote
}
type TICK struct{}
type DONE struct {
	WHO string
}
type OrderStatus int

const (
	Created OrderStatus = iota
	Placed
	PartiallyFilled
	Filled
	Canceled
)

type OrderType int

const (
	Bid OrderType = iota
	Ask
)

type Quote struct {
	LastPrice  float32
	CurrentBid float32
	CurrentAsk float32
}
type Order struct {
	Timestamp time.Time
	Quantity  int
	Price     float32
	Filled    int
	Status    OrderStatus
	Origin    actor.PID
	OrderType
	OrderID string
}
