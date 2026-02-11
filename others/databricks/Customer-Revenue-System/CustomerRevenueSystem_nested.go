package databricks

import "container/heap"

// ---------------------------------------------------------------------------
// Shared tree node used by both nested revenue strategies.
// ---------------------------------------------------------------------------

type nestedNode struct {
	ownRevenue int
	parentId   int   // -1 = root (no referrer)
	children   []int // child customer IDs
}

// topKByRevenue returns up to k IDs whose revenues[id] >= minRevenue,
// sorted descending by revenue.  Reuses the minHeap from the base file.
func topKByRevenue(revenues []int, k, minRevenue int) []int {
	h := &minHeap{}
	for i, rev := range revenues {
		if rev < minRevenue {
			continue
		}
		c := &customer{id: i, totalRevenue: rev}
		if h.Len() < k {
			heap.Push(h, c)
			continue
		}
		if (*h)[0].totalRevenue < rev {
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

// ===========================================================================
// Read-Heavy: eager propagation on write
//
//   AddByReferral  O(D)       — walk up parent chain, add revenue to each ancestor
//   GetTopKCustomer O(N log K) — use precomputed nestedRevenue directly
//
// Best when reads (GetTopKCustomer) are much more frequent than writes.
// ===========================================================================

// NestedRevenueReadHeavy eagerly maintains every customer's nested revenue
// (own + entire subtree) by propagating up the parent chain on each write.
type NestedRevenueReadHeavy struct {
	nodes         []*nestedNode
	nestedRevenue []int // always up-to-date
}

func NewNestedRevenueReadHeavy() *NestedRevenueReadHeavy {
	return &NestedRevenueReadHeavy{}
}

func (rs *NestedRevenueReadHeavy) Add(revenue int) int {
	id := len(rs.nodes)
	rs.nodes = append(rs.nodes, &nestedNode{ownRevenue: revenue, parentId: -1})
	rs.nestedRevenue = append(rs.nestedRevenue, revenue)
	return id
}

func (rs *NestedRevenueReadHeavy) AddByReferral(revenue, referrerId int) int {
	if referrerId < 0 || referrerId >= len(rs.nodes) {
		return -1
	}
	id := len(rs.nodes)
	rs.nodes = append(rs.nodes, &nestedNode{ownRevenue: revenue, parentId: referrerId})
	rs.nestedRevenue = append(rs.nestedRevenue, revenue)
	rs.nodes[referrerId].children = append(rs.nodes[referrerId].children, id)
	// Propagate revenue up to all ancestors — O(D).
	for cur := referrerId; cur != -1; cur = rs.nodes[cur].parentId {
		rs.nestedRevenue[cur] += revenue
	}
	return id
}

func (rs *NestedRevenueReadHeavy) GetTopKCustomer(k, minRevenue int) []int {
	return topKByRevenue(rs.nestedRevenue, k, minRevenue)
}

// ===========================================================================
// Write-Heavy: lazy aggregation on read
//
//   AddByReferral  O(1)       — store parent pointer only
//   GetTopKCustomer O(N log K) — bottom-up aggregation then top-K heap
//
// Best when writes (AddByReferral) are much more frequent than reads.
// ===========================================================================

// NestedRevenueWriteHeavy stores only the tree structure and recomputes
// nested revenue on demand via a single bottom-up pass.
type NestedRevenueWriteHeavy struct {
	nodes []*nestedNode
}

func NewNestedRevenueWriteHeavy() *NestedRevenueWriteHeavy {
	return &NestedRevenueWriteHeavy{}
}

func (rs *NestedRevenueWriteHeavy) Add(revenue int) int {
	id := len(rs.nodes)
	rs.nodes = append(rs.nodes, &nestedNode{ownRevenue: revenue, parentId: -1})
	return id
}

func (rs *NestedRevenueWriteHeavy) AddByReferral(revenue, referrerId int) int {
	if referrerId < 0 || referrerId >= len(rs.nodes) {
		return -1
	}
	id := len(rs.nodes)
	rs.nodes = append(rs.nodes, &nestedNode{ownRevenue: revenue, parentId: referrerId})
	rs.nodes[referrerId].children = append(rs.nodes[referrerId].children, id)
	return id // O(1) — no propagation
}

func (rs *NestedRevenueWriteHeavy) GetTopKCustomer(k, minRevenue int) []int {
	n := len(rs.nodes)
	nested := make([]int, n)
	// Bottom-up: children always have higher IDs than parents (auto-increment).
	for i := n - 1; i >= 0; i-- {
		nested[i] = rs.nodes[i].ownRevenue
		for _, childId := range rs.nodes[i].children {
			nested[i] += nested[childId]
		}
	}
	return topKByRevenue(nested, k, minRevenue)
}
