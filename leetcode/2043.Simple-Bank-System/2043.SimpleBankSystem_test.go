package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleBankSystem(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	assert.True(t, bank.Withdraw(3, 10))
	assert.True(t, bank.Transfer(5, 1, 20))
	assert.True(t, bank.Deposit(5, 20))
	assert.False(t, bank.Transfer(3, 4, 15))
	assert.False(t, bank.Withdraw(10, 50))
}
