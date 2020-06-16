package main

import "time"

import "marketplace/common"

type SubmitOrderRequest struct {
	Stock       string
	OrderDetail common.Order
}
type OrderConfirmation struct {
}
type OrderFulfillment struct {
	Timestamp time.Time
	Trades    []common.Trade
	OrderID   string
}
type GetQuoteMessage struct {
	Stock string
}
type OrderStatus int

const (
	Created OrderStatus = iota
	Placed
	PartiallyFilled
	Filled
	Canceled
)
