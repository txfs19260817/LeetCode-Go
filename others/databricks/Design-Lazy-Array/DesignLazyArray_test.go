package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyArray_Example1(t *testing.T) {
	la := NewLazyArray([]int{10, 20, 30, 40, 50})
	doubled := la.Map(func(n int) int { return n * 2 })
	assert.Equal(t, 1, doubled.IndexOf(40))
}

func TestLazyArray_Example2(t *testing.T) {
	la := NewLazyArray([]int{10, 20, 30, 40, 50})
	doubled := la.Map(func(n int) int { return n * 2 })
	tripled := doubled.Map(func(n int) int { return n * 3 })
	assert.Equal(t, 3, tripled.IndexOf(240))
}

func TestLazyArray_Example3_NotFound(t *testing.T) {
	la := NewLazyArray([]int{1, 2, 3, 4, 5})
	mapped := la.Map(func(n int) int { return n + 10 })
	assert.Equal(t, -1, mapped.IndexOf(100))
}

func TestLazyArray_ChainOfThreeMaps(t *testing.T) {
	la := NewLazyArray([]int{1, 2, 3, 4, 5})
	// 1→(+1)=2→(*3)=6→(-1)=5   index 0
	// 2→(+1)=3→(*3)=9→(-1)=8   index 1
	// 3→(+1)=4→(*3)=12→(-1)=11 index 2
	// 4→(+1)=5→(*3)=15→(-1)=14 index 3
	// 5→(+1)=6→(*3)=18→(-1)=17 index 4
	result := la.
		Map(func(n int) int { return n + 1 }).
		Map(func(n int) int { return n * 3 }).
		Map(func(n int) int { return n - 1 })
	assert.Equal(t, 0, result.IndexOf(5))
	assert.Equal(t, 2, result.IndexOf(11))
	assert.Equal(t, 4, result.IndexOf(17))
	assert.Equal(t, -1, result.IndexOf(99))
}

func TestLazyArray_SingleElement(t *testing.T) {
	la := NewLazyArray([]int{42})
	mapped := la.Map(func(n int) int { return n * 2 })
	assert.Equal(t, 0, mapped.IndexOf(84))
	assert.Equal(t, -1, mapped.IndexOf(42))
}

func TestLazyArray_OriginalUnchanged(t *testing.T) {
	la := NewLazyArray([]int{10, 20, 30})
	_ = la.Map(func(n int) int { return n * 100 })
	// Original should still find raw values
	assert.Equal(t, 1, la.IndexOf(20))
}

func TestLazyArray_FilterOnly(t *testing.T) {
	la := NewLazyArray([]int{10, 20, 30, 40, 50})
	filtered := la.Filter(func(n int) bool { return n%20 == 0 })
	assert.Equal(t, 1, filtered.IndexOf(20))
	assert.Equal(t, 3, filtered.IndexOf(40))
	assert.Equal(t, -1, filtered.IndexOf(10))
}

func TestLazyArray_FilterAfterMap(t *testing.T) {
	la := NewLazyArray([]int{1, 2, 3, 4, 5})
	result := la.
		Map(func(n int) int { return n * 3 }).
		Filter(func(n int) bool { return n%2 == 0 })
	// mapped -> [3,6,9,12,15], filtered -> [6,12]
	assert.Equal(t, 1, result.IndexOf(6))
	assert.Equal(t, 3, result.IndexOf(12))
	assert.Equal(t, -1, result.IndexOf(9))
}

func TestLazyArray_FilterBeforeMap(t *testing.T) {
	la := NewLazyArray([]int{1, 2, 3, 4, 5})
	result := la.
		Filter(func(n int) bool { return n%2 == 1 }).
		Map(func(n int) int { return n * 10 })
	// filtered odds -> [1,3,5], mapped -> [10,30,50]
	assert.Equal(t, 0, result.IndexOf(10))
	assert.Equal(t, 4, result.IndexOf(50))
	assert.Equal(t, -1, result.IndexOf(20))
}

func TestLazyArray_EarlyTermination_AvoidsLaterPanic(t *testing.T) {
	la := NewLazyArray([]int{3, 4, 5, 6})
	processed := make([]int, 0, 2)
	mapped := la.
		Filter(func(n int) bool { return n%2 == 1 }).
		Map(func(n int) int {
			if n > 5 {
				panic("should not reach value 6")
			}
			processed = append(processed, n)
			return n
		})

	assert.NotPanics(t, func() {
		assert.Equal(t, 2, mapped.IndexOf(5))
	})
	assert.Equal(t, []int{3, 5}, processed)
}
