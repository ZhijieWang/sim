package market

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type Quote struct {
	LastTradedPrice  float32
	LastTradedVolume int
	CurrentBid       float32
	BidVolumer       int
	CurrentAsk       float32
	AskVolume        int
}

const (
	// MaxPriceRange bounds the order book price point size.
	// If an order's price is within the current price +/- max price range,
	// the order will be condiser as a valid one. Or it will be rejected.
	MaxPriceRange = 10
)

type Market interface {
	//OrderHandler(Order) bool
	GetSymbol() string
	//	MatchOrder(Order) []Trade
	GetQuote() Quote
	init(float32, int)
}

func (m *marketImpl) GetQuote() Quote {
	return Quote{
		m.Price,
		m.Orders.Entries[m.Orders.bidMax].Price,
		m.Orders.Entries[m.Orders.askMin].Price,
	}
}

type TradeType int

const (
	Bought TradeType = iota
	Sold
)

// Trade defines the message of a trade to be placed
type Trade struct {
	Stock        string
	Position     int
	AveragePrice float32
	TradeType    TradeType
}
