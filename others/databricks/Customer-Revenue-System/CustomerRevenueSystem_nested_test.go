package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Shared test scenario:
//
//	0 (own=100)
//	└── 2 (own=40)
//	    └── 3 (own=25)
//	1 (own=60)
//	└── 4 (own=10)
//
// Nested revenues: 0→165, 1→70, 2→65, 3→25, 4→10

// --------------- Read-Heavy ---------------

func TestNestedReadHeavy_Basic(t *testing.T) {
	rs := NewNestedRevenueReadHeavy()

	assert.Equal(t, 0, rs.Add(100))
	assert.Equal(t, 1, rs.Add(60))
	assert.Equal(t, 2, rs.AddByReferral(40, 0)) // 0→2
	assert.Equal(t, 3, rs.AddByReferral(25, 2)) // 0→2→3
	assert.Equal(t, 4, rs.AddByReferral(10, 1)) // 1→4

	assert.Equal(t, []int{0, 1, 2}, rs.GetTopKCustomer(3, 0))
	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 50))
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(5, 100))
}

func TestNestedReadHeavy_DeepChain(t *testing.T) {
	rs := NewNestedRevenueReadHeavy()

	assert.Equal(t, 0, rs.Add(10))
	assert.Equal(t, 1, rs.AddByReferral(20, 0)) // 0→1
	assert.Equal(t, 2, rs.AddByReferral(30, 1)) // 0→1→2
	assert.Equal(t, 3, rs.AddByReferral(40, 2)) // 0→1→2→3

	// Nested: 0=10+20+30+40=100, 1=20+30+40=90, 2=30+40=70, 3=40
	assert.Equal(t, []int{0, 1, 2, 3}, rs.GetTopKCustomer(10, 0))
	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 0))
}

func TestNestedReadHeavy_InvalidReferrer(t *testing.T) {
	rs := NewNestedRevenueReadHeavy()
	assert.Equal(t, 0, rs.Add(100))
	assert.Equal(t, -1, rs.AddByReferral(50, 5))
	assert.Equal(t, -1, rs.AddByReferral(50, -1))
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(10, 0))
}

// --------------- Write-Heavy ---------------

func TestNestedWriteHeavy_Basic(t *testing.T) {
	rs := NewNestedRevenueWriteHeavy()

	assert.Equal(t, 0, rs.Add(100))
	assert.Equal(t, 1, rs.Add(60))
	assert.Equal(t, 2, rs.AddByReferral(40, 0))
	assert.Equal(t, 3, rs.AddByReferral(25, 2))
	assert.Equal(t, 4, rs.AddByReferral(10, 1))

	assert.Equal(t, []int{0, 1, 2}, rs.GetTopKCustomer(3, 0))
	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 50))
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(5, 100))
}

func TestNestedWriteHeavy_DeepChain(t *testing.T) {
	rs := NewNestedRevenueWriteHeavy()

	assert.Equal(t, 0, rs.Add(10))
	assert.Equal(t, 1, rs.AddByReferral(20, 0))
	assert.Equal(t, 2, rs.AddByReferral(30, 1))
	assert.Equal(t, 3, rs.AddByReferral(40, 2))

	assert.Equal(t, []int{0, 1, 2, 3}, rs.GetTopKCustomer(10, 0))
	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 0))
}

func TestNestedWriteHeavy_InvalidReferrer(t *testing.T) {
	rs := NewNestedRevenueWriteHeavy()
	assert.Equal(t, 0, rs.Add(100))
	assert.Equal(t, -1, rs.AddByReferral(50, 5))
	assert.Equal(t, -1, rs.AddByReferral(50, -1))
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(10, 0))
}

// --------------- Both strategies produce identical results ---------------

func TestNestedStrategies_SameResults(t *testing.T) {
	rh := NewNestedRevenueReadHeavy()
	wh := NewNestedRevenueWriteHeavy()

	ops := []struct {
		revenue    int
		referrerId int // -2 = use Add instead of AddByReferral
	}{
		{100, -2}, {60, -2}, {40, 0}, {25, 2}, {10, 1},
		{15, 3}, {5, 0}, {35, 4},
	}
	for _, op := range ops {
		if op.referrerId == -2 {
			rh.Add(op.revenue)
			wh.Add(op.revenue)
		} else {
			rh.AddByReferral(op.revenue, op.referrerId)
			wh.AddByReferral(op.revenue, op.referrerId)
		}
	}

	for _, k := range []int{1, 3, 5, 10} {
		for _, minRev := range []int{0, 50, 100, 200} {
			assert.Equal(t, rh.GetTopKCustomer(k, minRev),
				wh.GetTopKCustomer(k, minRev),
				"k=%d, minRev=%d", k, minRev)
		}
	}
}
