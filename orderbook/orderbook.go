package orderbook

import (
	"math"
)

type TradeStatus string
type OrderType int

const (
	BidOrder OrderType = iota
	AskOrder
)

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

func NewOrderBook() OrderBook {
	o := &orderBookImpl{
		bids: []*entry{},
		asks: []*entry{},
	}
	return o
}
func (o *orderBookImpl) PlaceBid(order Order) (bool, []Trade) {
	// A bid that is at or above currentAsk, should be matched
	// Otherwise, insert into the list of entries

	currentAsk, _ := o.GetCurrentAsk()
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

// GetCurrentAsk return the information on current/lowest ask available.
// if no ask order in the system, return positive infinity for price and 0 for quantity
func (o *orderBookImpl) GetCurrentAsk() (float64, int) {

	if len(o.asks) == 0 {
		return math.Inf(1), 0
	} else {
		cur := o.asks[0]
		return cur.GetPrice(), cur.GetVolume()
	}
}

// GetCurrentBid return the information on current/highest bid availablke.
// if no bid order in the system, return 0 for price and 0 for quantity
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
