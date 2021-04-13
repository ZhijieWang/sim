package common_test

import (
	"marketplace/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountCopy(t *testing.T) {

	base := common.NewDefaultAccount(1000000)
	target := base.Copy()
	target.Commit("A", 100, 1.0, common.BidOrder)
	assert.Equal(t, base.GetId(), target.GetId())
	assert.NotEqual(t, base, target)
	assert.NotEqual(t, base.GetBalance(), target.GetBalance())
}
func TestPositionCopy(t *testing.T) {
	var base, target common.Position
	base = common.NewCashPosition(100.0)
	target = base.Copy()
	target.Commit(100, 1.0, "a")
	assert.Equal(t, base.GetSymbol(), target.GetSymbol())
	assert.NotEqual(t, base.GetAvailable(), target.GetAvailable())
	assert.NotEqual(t, base, target)
	assert.NotEqual(t, base.GetBalance(), target.GetBalance())
}
func TestAccount(t *testing.T) {
	act := common.NewDefaultAccount(100000)
	assert.NotZero(t, act.GetId())
	assert.NotNil(t, act.GetBalance()["$"])
}
func TestAccountOrder(t *testing.T) {
	order, _ := common.NewDefaultAccount(10000).Commit("A", 100, 1, common.BidOrder)
	assert.NotNil(t, order)
	assert.NotZero(t, order.GetId().AccountId)
}

func TestSetOrderAccountId(t *testing.T) {
	order := common.NewOrder(common.BidOrder, 10, 1, "A")
	assert.Zero(t, order.GetId().AccountId)
	order.SetTraderId(1)
	assert.NotZero(t, order.GetId().AccountId)
}
