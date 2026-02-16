package databricks

// LazyArray supports deferred mapping operations on an integer array.
type LazyArray struct {
	arr []int
	ops []opFn
}

// opFn is a unified deferred operation.
// It returns the next value and whether the element should remain in the pipeline.
type opFn func(int) (next int, keep bool)

// NewLazyArray initializes a LazyArray with the given slice.
func NewLazyArray(arr []int) *LazyArray {
	return &LazyArray{arr: arr}
}

// Map returns a new LazyArray that shares the same underlying data
// but has fn appended to its transformation pipeline.
// Time: O(k) to copy existing pipeline of length k.
// Extra space: O(k) for the copied pipeline.
func (la *LazyArray) Map(fn func(int) int) *LazyArray {
	copied := make([]opFn, len(la.ops), len(la.ops)+1)
	copy(copied, la.ops)
	copied = append(copied, func(v int) (int, bool) {
		return fn(v), true
	})
	return &LazyArray{arr: la.arr, ops: copied}
}

// Filter returns a new LazyArray that shares the same underlying data
// but has predicate appended to its pipeline. Elements that don't satisfy
// the predicate are skipped.
// Time: O(k) to copy existing pipeline of length k.
// Extra space: O(k) for the copied pipeline.
func (la *LazyArray) Filter(predicate func(int) bool) *LazyArray {
	copied := make([]opFn, len(la.ops), len(la.ops)+1)
	copy(copied, la.ops)
	copied = append(copied, func(v int) (int, bool) {
		return v, predicate(v)
	})
	return &LazyArray{arr: la.arr, ops: copied}
}

// IndexOf applies all accumulated functions to each original element
// and returns the index of the first result equal to target, or -1.
// Time: O(n*k) in the worst case where n is array length and k is pipeline length.
// Extra space: O(1) (not counting the already-built pipeline).
func (la *LazyArray) IndexOf(target int) int {
	for i, v := range la.arr {
		result := v
		kept := true
		for _, op := range la.ops {
			next, ok := op(result)
			if !ok {
				kept = false
				break
			}
			result = next
		}
		if kept && result == target {
			return i
		}
	}
	return -1
}
