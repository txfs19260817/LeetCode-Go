package databricks

import "container/heap"

type customer struct {
	id           int
	totalRevenue int // own revenue + revenue from direct referrals
}

// minHeap is a min-heap of customers ordered by totalRevenue.
type minHeap []*customer

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].totalRevenue < h[j].totalRevenue }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(*customer)) }
func (h *minHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// RevenueSystem tracks customers and their revenue.
type RevenueSystem struct {
	customers []*customer
}

// NewRevenueSystem initializes the system.
func NewRevenueSystem() *RevenueSystem {
	return &RevenueSystem{}
}

// Add adds a new customer with the given revenue. Returns the auto-incremented ID.
func (rs *RevenueSystem) Add(revenue int) int {
	id := len(rs.customers)
	rs.customers = append(rs.customers, &customer{id: id, totalRevenue: revenue})
	return id
}

// AddByReferral adds a new customer referred by referrerId.
// The referrer's total revenue increases by revenue.
// Returns -1 if referrerId is invalid.
func (rs *RevenueSystem) AddByReferral(revenue int, referrerId int) int {
	if referrerId < 0 || referrerId >= len(rs.customers) {
		return -1
	}
	id := rs.Add(revenue)
	rs.customers[referrerId].totalRevenue += revenue
	return id
}

// GetTopKCustomer returns up to k customer IDs with totalRevenue >= minRevenue,
// sorted descending by total revenue. Uses a min-heap of size k.
func (rs *RevenueSystem) GetTopKCustomer(k, minRevenue int) []int {
	h := &minHeap{}
	for _, c := range rs.customers {
		if c.totalRevenue < minRevenue {
			continue
		}
		if h.Len() < k {
			heap.Push(h, c)
			continue
		}
		if (*h)[0].totalRevenue < c.totalRevenue {
			(*h)[0] = c
			heap.Fix(h, 0)
		}
	}
	result := make([]int, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(*customer).id
	}
	return result
}

// getTopKCustomerPushPop keeps the original push-then-pop approach for benchmarking.
func (rs *RevenueSystem) getTopKCustomerPushPop(k, minRevenue int) []int {
	h := &minHeap{}
	for _, c := range rs.customers {
		if c.totalRevenue >= minRevenue {
			heap.Push(h, c)
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	result := make([]int, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(*customer).id
	}
	return result
}
