package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountCopy(t *testing.T) {

	base := NewDefaultAccount(1000000)
	target := base.Copy()
	target.Commit("A", 100, 1.0, BidOrder)
	assert.Equal(t, base.GetId(), target.GetId())
	assert.NotEqual(t, base, target)
	assert.NotEqual(t, base.GetBalance(), target.GetBalance())
}
func TestPositionCopy(t *testing.T) {
	var base, target Position
	base = NewCashPosition(100.0)
	target = base.Copy()
	target.Commit(100, 1.0, "a")
	assert.Equal(t, base.GetSymbol(), target.GetSymbol())
	assert.NotEqual(t, base.GetAvailable(), target.GetAvailable())
	assert.NotEqual(t, base, target)
	assert.NotEqual(t, base.GetBalance(), target.GetBalance())
}
