package _341_Flatten_Nested_List_Iterator

// NestedInteger This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return false }

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0 }

// SetInteger Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }

// Solution

type NestedIterator struct {
	arr []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var arr []int
	var dfs func([]*NestedInteger)
	dfs = func(nestedList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() {
				arr = append(arr, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{arr}
}

func (it *NestedIterator) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *NestedIterator) HasNext() bool {
	return len(it.arr) > 0
}
