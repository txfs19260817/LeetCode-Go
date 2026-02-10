package databricks

import (
	"math/rand"
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

func BenchmarkGetTopKCustomer_Scenarios(b *testing.B) {
	const n = 100000
	uniform := makeUniformRevenues(n, 1000, 1)
	increasing := makeIncreasingRevenues(n)
	methods := []struct {
		name string
		fn   func(*RevenueSystem, int, int) []int
	}{
		{name: "PushPop", fn: (*RevenueSystem).getTopKCustomerPushPop},
		{name: "TopCompare", fn: (*RevenueSystem).GetTopKCustomer},
	}

	cases := []struct {
		name       string
		revenues   []int
		k          int
		minRevenue int
	}{
		{name: "SparseEligible_SmallK", revenues: uniform, k: 10, minRevenue: 950},
		{name: "SparseEligible_LargeK", revenues: uniform, k: 20000, minRevenue: 950},
		{name: "DenseEligible_SmallK", revenues: uniform, k: 10, minRevenue: 100},
		{name: "DenseEligible_LargeK", revenues: uniform, k: 50000, minRevenue: 100},
		{name: "AllEligible_WorstCaseInc", revenues: increasing, k: 10, minRevenue: 0},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for _, method := range methods {
				b.Run(method.name, func(b *testing.B) {
					benchmarkGetTopKCustomer(b, method.fn, tc.revenues, tc.k, tc.minRevenue)
				})
			}
		})
	}
}

func benchmarkGetTopKCustomer(b *testing.B, fn func(*RevenueSystem, int, int) []int, revenues []int, k, minRevenue int) {
	rs := buildRevenueSystem(revenues)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = fn(rs, k, minRevenue)
	}
}

func buildRevenueSystem(revenues []int) *RevenueSystem {
	rs := NewRevenueSystem()
	for _, revenue := range revenues {
		rs.Add(revenue)
	}
	return rs
}

func makeUniformRevenues(n, max int, seed int64) []int {
	rnd := rand.New(rand.NewSource(seed))
	revenues := make([]int, n)
	for i := range revenues {
		revenues[i] = rnd.Intn(max)
	}
	return revenues
}

func makeIncreasingRevenues(n int) []int {
	revenues := make([]int, n)
	for i := range revenues {
		revenues[i] = i
	}
	return revenues
}
