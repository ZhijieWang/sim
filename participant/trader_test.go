package participant_test

import (
	"marketplace/common"
	"marketplace/participant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketMakerFunc(t *testing.T) {
	maker := participant.NewMarketMaker("A")
	o := maker.Trade(nil)
	assert.NotNil(t, o)
}
func TestTraderFuncWithoutQuote(t *testing.T) {
	trader := participant.NewParticipant(1, 1000.0)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should have panicked!")
		}
	}()
	trader.Trade(nil)
}
func TestTraderFuncWithQuote(t *testing.T) {
	trader := participant.NewParticipant(1, 1000.0)
	o := trader.Trade(map[string]common.Quote{
		"A": {
			LastTradedPrice:  1.0,
			LastTradedVolume: 1.0,
			CurrentAsk:       1.0,
			CurrentBid:       1.0,
		},
	})
	assert.NotNil(t, o)
}
