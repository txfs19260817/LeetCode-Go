package databricks

// LazyArray supports deferred mapping operations on an integer array.
type LazyArray struct {
	arr []int
	fns []func(int) int
}

// NewLazyArray initializes a LazyArray with the given slice.
func NewLazyArray(arr []int) *LazyArray {
	return &LazyArray{arr: arr}
}

// Map returns a new LazyArray that shares the same underlying data
// but has fn appended to its transformation pipeline.
func (la *LazyArray) Map(fn func(int) int) *LazyArray {
	copied := make([]func(int) int, len(la.fns), len(la.fns)+1)
	copy(copied, la.fns)
	copied = append(copied, fn)
	return &LazyArray{arr: la.arr, fns: copied}
}

// IndexOf applies all accumulated functions to each original element
// and returns the index of the first result equal to target, or -1.
func (la *LazyArray) IndexOf(target int) int {
	for i, v := range la.arr {
		result := v
		for _, fn := range la.fns {
			result = fn(result)
		}
		if result == target {
			return i
		}
	}
	return -1
}
