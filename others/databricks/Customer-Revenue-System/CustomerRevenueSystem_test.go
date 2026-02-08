package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevenueSystem_MainExample(t *testing.T) {
	rs := NewRevenueSystem()

	assert.Equal(t, 0, rs.Add(100))             // Customer 0: $100
	assert.Equal(t, 1, rs.Add(50))              // Customer 1: $50
	assert.Equal(t, 2, rs.AddByReferral(30, 0)) // Customer 0: $130, Customer 2: $30
	assert.Equal(t, 3, rs.AddByReferral(70, 1)) // Customer 1: $120, Customer 3: $70
	assert.Equal(t, 4, rs.Add(50))              // Customer 4: $50

	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 100))

	assert.Equal(t, 5, rs.AddByReferral(50, 4)) // Customer 4: $100, Customer 5: $50

	assert.Equal(t, []int{0, 1, 4}, rs.GetTopKCustomer(3, 100))
}

func TestRevenueSystem_InvalidReferrer(t *testing.T) {
	rs := NewRevenueSystem()

	assert.Equal(t, 0, rs.Add(100))

	// Referrer ID 5 does not exist
	assert.Equal(t, -1, rs.AddByReferral(50, 5))
	// Negative referrer ID
	assert.Equal(t, -1, rs.AddByReferral(50, -1))

	// Only customer 0 should exist
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(10, 0))
}

func TestRevenueSystem_MultipleReferrals(t *testing.T) {
	rs := NewRevenueSystem()

	assert.Equal(t, 0, rs.Add(100))             // Customer 0: $100
	assert.Equal(t, 1, rs.AddByReferral(50, 0)) // Customer 0: $150, Customer 1: $50
	assert.Equal(t, 2, rs.AddByReferral(30, 0)) // Customer 0: $180, Customer 2: $30
	assert.Equal(t, 3, rs.AddByReferral(20, 0)) // Customer 0: $200, Customer 3: $20

	// Only customer 0 has >= 100
	assert.Equal(t, []int{0}, rs.GetTopKCustomer(5, 100))
	// All customers with >= 0
	assert.Equal(t, []int{0, 1, 2, 3}, rs.GetTopKCustomer(10, 0))
	// Top 2 with >= 0
	assert.Equal(t, []int{0, 1}, rs.GetTopKCustomer(2, 0))
}
