package leetcode

// NestedInteger This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	isInteger bool
	value     int
	list      []*NestedInteger
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return n.isInteger }

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return n.value }

// SetInteger Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.isInteger = true
	n.value = value
	n.list = nil
}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	if n.isInteger {
		n.isInteger = false
		n.value = 0
	}
	n.list = append(n.list, &elem)
}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return n.list }

// Solution

type NestedIterator struct {
	stack []frame
}

// frame records the current scan position of one nested list level.
type frame struct {
	list []*NestedInteger
	idx  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []frame{{list: nestedList, idx: 0}}}
}

func (it *NestedIterator) Next() int {
	// HasNext expands nested lists until the stack top points to an integer.
	it.HasNext()
	top := &it.stack[len(it.stack)-1]
	val := top.list[top.idx].GetInteger()
	top.idx++
	return val
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		top := &it.stack[len(it.stack)-1]

		// This level is exhausted, move back to its parent level.
		if top.idx == len(top.list) {
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}

		cur := top.list[top.idx]
		if cur.IsInteger() {
			return true
		}

		// Expand one nested list lazily instead of flattening all values up front.
		top.idx++
		it.stack = append(it.stack, frame{list: cur.GetList()})
	}

	return false
}

// NestedIterator2 eagerly flattens all integers during construction.
type NestedIterator2 struct {
	arr []int
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator2 {
	var arr []int
	var dfs func([]*NestedInteger)
	dfs = func(list []*NestedInteger) {
		for _, nest := range list {
			if nest.IsInteger() {
				arr = append(arr, nest.GetInteger())
				continue
			}
			dfs(nest.GetList())
		}
	}
	dfs(nestedList)
	return &NestedIterator2{arr: arr}
}

func (it *NestedIterator2) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *NestedIterator2) HasNext() bool {
	return len(it.arr) > 0
}
