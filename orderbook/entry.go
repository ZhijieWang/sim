package orderbook

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
