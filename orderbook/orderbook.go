package orderbook

import (
	"math/rand"
)

type Order interface {
	GetType() OrderType
	GetPrice() float32
	GetVolume() int
}

type TradeStatus string
type OrderType int

const (
	BidOrder OrderType = iota
	AskOrder
)

type Trade interface {
	GetStatus() TradeStatus
	GetFilledQuantity() int
	GetFilledAvgPrice() int
	GetOrderId()
}

func NewOrder(o OrderType, price float32, v int) Order {
	return &orderImpl{
		ID:       rand.Int(),
		BidOrAsk: o,
		Price:    price,
		Volume:   v,
	}
}

type orderImpl struct {
	ID       int
	Ticker   string
	BidOrAsk OrderType // true for ask, false for bid
	Price    float32
	Volume   int
}

func (o *orderImpl) GetPrice() float32 {
	return o.Price
}
func (o *orderImpl) GetVolume() int {
	return o.Volume
}

type entry struct {
	orders []Order
	volume int
}

func (e *entry) Add(o Order) {

}

type OrderBook interface {
	PlaceBid(Order) bool
	PlaceAsk(Order) bool
	GetCurrentBid() (float32, int)
	GetCurrentAsk() (float32, int)
	Fill(Order) []Trade
}
type orderBookImpl struct {
	currentBid float32
	currentAsk float32
	bids       []Order // sorted in dessending orders
	asks       []Order // sorted in ascending orders
}

func NewOrderBook(initPrice float32, volume int) OrderBook {
	return &orderBookImpl{
		bids: []Order{NewOrder(AskOrder, initPrice, volume)},
		asks: []Order{},
	}
}
func (o *orderBookImpl) PlaceBid(Order) bool {
	return true
}
func (o *orderBookImpl) PlaceAsk(Order) bool {
	return true
}

func (o *orderBookImpl) GetCurrentBid() (float32, int) {
	cur := o.bids[0]
	return cur.GetPrice(), cur.GetVolume()
}
func (o *orderBookImpl) GetCurrentAsk() (float32, int) {
	cur := o.asks[0]
	return cur.GetPrice(), cur.GetVolume()
}
func (o *orderBookImpl) Fill(order Order) []Trade {
	var trades []Trade
	switch order.GetType() {
	case BidOrder:
		if o.currentAsk < order.GetPrice() {
			var i int = 0
			q = order.GetVolume()
			for order.bids[i].currentPrice()&q > 0 {

			}
		}

	case AskOrder:
	}

}
