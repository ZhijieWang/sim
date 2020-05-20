package orderbook

import (
	"marketplace/common"
	"math"
)

func NewOrderBook() OrderBook {
	o := &orderBookImpl{
		bids: []*entry{},
		asks: []*entry{},
	}
	return o
}

type OrderBook interface {
	PlaceBid(common.Order) (bool, []common.Trade)
	PlaceAsk(common.Order) (bool, []common.Trade)
	GetCurrentBid() (float64, int)
	GetCurrentAsk() (float64, int)
	Fill(common.Order) []common.Trade
}
type orderBookImpl struct {
	bids []*entry // sorted in descending common.Orders
	asks []*entry // sorted in ascending common.Orders
}

func (o *orderBookImpl) PlaceBid(order common.Order) (bool, []common.Trade) {
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
func InsertEntry(book *[]*entry, order common.Order) {
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

func (o *orderBookImpl) PlaceAsk(order common.Order) (bool, []common.Trade) {
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
// if no ask common.Order in the system, return positive infinity for price and 0 for quantity
func (o *orderBookImpl) GetCurrentAsk() (float64, int) {

	if len(o.asks) == 0 {
		return math.Inf(1), 0
	} else {
		cur := o.asks[0]
		return cur.GetPrice(), cur.GetVolume()
	}
}

// GetCurrentBid return the information on current/highest bid availablke.
// if no bid common.Order in the system, return 0 for price and 0 for quantity
func (o *orderBookImpl) GetNextBidOrder() common.Order {
	if len(o.bids) != 0 {
		return o.bids[0].Peek()
	} else {
		return nil
	}
}
func (o *orderBookImpl) GetNextAskOrder() common.Order {
	if len(o.asks) != 0 {
		return o.asks[0].Peek()
	} else {
		return nil
	}
}

func (o *orderBookImpl) Fill(order common.Order) []common.Trade {
	var trades []common.Trade
	switch order.GetType() {
	case common.BidOrder:
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
	case common.AskOrder:
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
