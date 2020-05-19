package orderbook

import (
	"fmt"
	"math/rand"
)

type Order interface {
	GetType() OrderType
	GetPrice() float64
	GetVolume() int
	Fill(int) Trade
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
	GetFilledAvgPrice() float64
	GetOrderId() int
}

type tradeImpl struct {
	tradeId int
	orderId int
	filled  int
	price   float64
	status  TradeStatus
}

func (t tradeImpl) GetStatus() TradeStatus {
	return t.status
}

func (t tradeImpl) GetFilledQuantity() int {
	return t.filled
}
func (t tradeImpl) GetFilledAvgPrice() float64 {
	return t.price
}

func (t tradeImpl) GetOrderId() int {
	return t.orderId
}
func NewTrade(oId int, filled int, price float64) Trade {
	return tradeImpl{
		tradeId: rand.Int(),
		orderId: oId,
		filled:  filled,
		price:   price,
		status:  "Filled",
	}

}
func NewOrder(o OrderType, price float64, v int) Order {
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
	Price    float64
	Volume   int
}

func (o *orderImpl) GetPrice() float64 {
	return o.Price
}
func (o *orderImpl) Fill(quantity int) Trade {
	if quantity > o.Volume {
		panic(fmt.Errorf("Filling more than what the order has. Want %d, Has %d", quantity, o.Volume))
	}
	o.Volume -= quantity
	if o.Volume == 0 {
		//o.Close()
	}

	return NewTrade(o.ID, quantity, o.Price)
}
func (o *orderImpl) GetVolume() int {
	return o.Volume
}
func (o *orderImpl) GetType() OrderType {
	return o.BidOrAsk
}

type entry struct {
	orders []Order
	volume int
	price  float64
}

func NewEntry(price float64) *entry {
	return &entry{
		orders: []Order{},
		volume: 0,
		price:  price,
	}
}
func (e *entry) Add(o Order) {
	e.orders = append(e.orders, o)
	e.volume += o.GetVolume()
}
func (e *entry) Peek() Order {
	return e.orders[0]
}
func (e *entry) Pop() Order {
	var x Order
	x, e.orders = e.orders[len(e.orders)-1], e.orders[:len(e.orders)-1]
	e.volume -= x.GetVolume()
	return x
}
func (e *entry) GetVolume() int {
	return e.volume
}
func (e *entry) GetPrice() float64 {
	return e.price
}

func (e *entry) Fill(volume int) []Trade {
	var fill int
	var trades []Trade
	for (volume != 0) && (0 < len(e.orders)) {
		if volume > e.orders[0].GetVolume() {
			fill = e.orders[0].GetVolume()
		} else {
			fill = volume
		}
		volume -= fill
		e.volume -= fill
		trades = append(trades, e.orders[0].Fill(fill))
		if e.orders[0].GetVolume() == 0 {
			e.orders = e.orders[1:]
			// empty order is cleared at the Entry level Fill method
		}

	}
	return trades
}

type OrderBook interface {
	PlaceBid(Order) (bool, []Trade)
	PlaceAsk(Order) (bool, []Trade)
	GetCurrentBid() (float64, int)
	GetCurrentAsk() (float64, int)
	Fill(Order) []Trade
}
type orderBookImpl struct {
	bids []*entry // sorted in descending orders
	asks []*entry // sorted in ascending orders
}

func NewOrderBook(initPrice float64, volume int) OrderBook {
	o := &orderBookImpl{
		bids: []*entry{},
		asks: []*entry{},
	}
	o.asks = append(o.bids, NewEntry(1.0))
	o.asks[0].Add(NewOrder(AskOrder, initPrice, volume))
	return o
}
func (o *orderBookImpl) PlaceBid(order Order) (bool, []Trade) {
	// A bid that is at or above currentAsk, should be matched
	// Otherwise, insert into the list of entries
	currentAsk, _ := o.GetCurrentAsk()
	fmt.Println(currentAsk)
	if order.GetPrice() >= currentAsk {
		return true, o.Fill(order)
	} else {
		InsertEntry(&o.bids, order)
		return true, nil
	}
}
func InsertEntry(book *[]*entry, order Order) {
	var e *entry
	for i := 0; i < len(*book); i++ {
		if (*book)[i].price > order.GetPrice() {
			*book = append(*book, nil)
			copy((*book)[i+1:], (*book)[i:])
			e = NewEntry(order.GetPrice())
			e.Add(order)
			(*book)[i] = e
			return

		}
		if (*book)[i].price == order.GetPrice() {
			(*book)[i].Add(order)
			return
		}
	}
	e = NewEntry(order.GetPrice())
	e.Add(order)
	*book = append(*book, e)
}

func (o *orderBookImpl) PlaceAsk(order Order) (bool, []Trade) {
	// An ask that is at or below currentBid, should be matched
	// Otherwise, insert into the list of entires
	currentBid, _ := o.GetCurrentBid()
	if order.GetPrice() <= currentBid {

		return true, o.Fill(order)
	} else {
		InsertEntry(&o.asks, order)
		return true, nil
	}
}

func (o *orderBookImpl) GetCurrentBid() (float64, int) {
	if len(o.bids) == 0 {
		return 0.0, 0
	} else {
		cur := o.bids[0]
		return cur.GetPrice(), cur.GetVolume()
	}
}
func (o *orderBookImpl) GetCurrentAsk() (float64, int) {
	if len(o.asks) == 0 {
		return 0.0, 0
	} else {
		cur := o.asks[0]
		return cur.GetPrice(), cur.GetVolume()
	}
}
func (o *orderBookImpl) GetNextBidOrder() Order {
	if len(o.bids) != 0 {
		return o.bids[0].Peek()
	} else {
		return nil
	}
}
func (o *orderBookImpl) GetNextAskOrder() Order {
	if len(o.asks) != 0 {
		return o.asks[0].Peek()
	} else {
		return nil
	}
}

func (o *orderBookImpl) Fill(order Order) []Trade {
	var trades []Trade
	switch order.GetType() {
	case BidOrder:
		price, unit := o.GetCurrentAsk()

		for unit != 0 && price <= order.GetPrice() && order.GetVolume() != 0 {
			if order.GetVolume() <= unit {
				unit = order.GetVolume()
			}
			trades = append(trades, order.Fill(unit))
			trades = append(trades, o.asks[0].Fill(unit)...)
			if o.asks[0].GetVolume() == 0 {
				o.asks = o.asks[1:]
			}
			price, unit = o.GetCurrentAsk()
		}
	case AskOrder:
		price, unit := o.GetCurrentBid()
		for unit != 0 && price >= order.GetPrice() && order.GetVolume() != 0 {
			if order.GetVolume() <= unit {
				unit = order.GetVolume()

				trades = append(trades, order.Fill(unit))
				trades = append(trades, o.bids[0].Fill(unit)...)
				if o.bids[0].GetVolume() == 0 {
					o.bids = o.bids[1:]
				}
				price, unit = o.GetCurrentBid()
			}
		}
	}
	return trades
}
