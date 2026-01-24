package uber

import "container/heap"

// RevenueTracker manages customer revenues and referral bonuses.
type RevenueTracker struct {
	revenues []int
}

type intMinHeap []int

func (h intMinHeap) Len() int {
	return len(h)
}

func (h intMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// NewRevenueTracker creates a new RevenueTracker.
func NewRevenueTracker() *RevenueTracker {
	return &RevenueTracker{
		revenues: make([]int, 0),
	}
}

// Add creates a new customer with the given revenue and returns its ID.
func (rt *RevenueTracker) Add(revenue int) int {
	rt.revenues = append(rt.revenues, revenue)
	return len(rt.revenues)
}

// AddByReferral creates a new customer referred by referrerID and returns its ID.
// The referrer's revenue increases by revenue, and the new customer's revenue
// is initialized to the referrer's revenue before the increase.
func (rt *RevenueTracker) AddByReferral(revenue int, referrerID int) int {
	if referrerID <= 0 || referrerID > len(rt.revenues) {
		return 0
	}
	referrerIndex := referrerID - 1
	preRevenue := rt.revenues[referrerIndex]
	rt.revenues[referrerIndex] = preRevenue + revenue
	rt.revenues = append(rt.revenues, preRevenue)
	return len(rt.revenues)
}

// ShowRevenue returns the current revenue for the given customer ID.
func (rt *RevenueTracker) ShowRevenue(id int) int {
	if id <= 0 || id > len(rt.revenues) {
		return 0
	}
	return rt.revenues[id-1]
}

// TopSmallestKCustomer returns IDs of the k smallest revenues strictly greater than minRevenue.
// If fewer than k customers satisfy the condition, return all of them.
// If multiple customers have the same revenue, return smaller IDs first.
func (rt *RevenueTracker) TopSmallestKCustomer(minRevenue int, k int) []int {
	revenueToIDs := make(map[int][]int)
	minHeap := &intMinHeap{}
	heap.Init(minHeap)

	for i, rev := range rt.revenues {
		if rev > minRevenue {
			if _, ok := revenueToIDs[rev]; !ok {
				heap.Push(minHeap, rev)
			}
			revenueToIDs[rev] = append(revenueToIDs[rev], i+1)
		}
	}
	result := make([]int, 0, k)
	for minHeap.Len() > 0 && len(result) < k {
		rev := heap.Pop(minHeap).(int)
		ids := revenueToIDs[rev]
		remaining := k - len(result)
		result = append(result, ids[:min(len(ids), remaining)]...)
	}

	return result
}
