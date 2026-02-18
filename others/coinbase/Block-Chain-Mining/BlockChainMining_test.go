package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockChainMiningBaseGreedyVsOptimal(t *testing.T) {
	miner := NewBlockChainMining()
	txs := []Transaction{
		{ID: "a", Size: 6, Fee: 13},
		{ID: "b", Size: 5, Fee: 10},
		{ID: "c", Size: 5, Fee: 10},
	}

	greedy := miner.MaxFeeGreedy(10, txs)
	optimal := miner.MaxFeeOptimal(10, txs)

	assert.Equal(t, 13, greedy.TotalFee)
	assert.Equal(t, 6, greedy.UsedSize)
	assert.Equal(t, 20, optimal.TotalFee)
	assert.Equal(t, 10, optimal.UsedSize)
}

func TestBlockChainMiningFollowUpBranchExclusive(t *testing.T) {
	miner := NewBlockChainMining()
	txs := []Transaction{
		{ID: "1", ParentID: "", Size: 2, Fee: 3},
		{ID: "2", ParentID: "1", Size: 2, Fee: 4},
		{ID: "3", ParentID: "2", Size: 3, Fee: 8},
		{ID: "4", ParentID: "2", Size: 3, Fee: 7},
		{ID: "5", ParentID: "", Size: 4, Fee: 9},
	}

	plan := miner.MaxFeeWithParents(14, txs)

	assert.Equal(t, 24, plan.TotalFee)
	assert.Equal(t, 11, plan.UsedSize)
	assert.Contains(t, plan.TxIDs, "5")
	has3 := contains(plan.TxIDs, "3")
	has4 := contains(plan.TxIDs, "4")
	assert.True(t, has3 != has4, "must not mine both sibling branches")
}

func TestBlockChainMiningFollowUpMultipleGroups(t *testing.T) {
	miner := NewBlockChainMining()
	txs := []Transaction{
		{ID: "A1", ParentID: "", Size: 2, Fee: 3},
		{ID: "A2", ParentID: "A1", Size: 3, Fee: 10},
		{ID: "A3", ParentID: "A1", Size: 3, Fee: 8},
		{ID: "B1", ParentID: "", Size: 3, Fee: 5},
		{ID: "B2", ParentID: "B1", Size: 2, Fee: 8},
	}

	plan := miner.MaxFeeWithParents(10, txs)
	assert.Equal(t, 26, plan.TotalFee)
	assert.Equal(t, 10, plan.UsedSize)
	assert.True(t, contains(plan.TxIDs, "A1"))
	assert.True(t, contains(plan.TxIDs, "A2"))
	assert.True(t, contains(plan.TxIDs, "B1"))
	assert.True(t, contains(plan.TxIDs, "B2"))
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
