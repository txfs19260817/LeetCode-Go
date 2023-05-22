package leetcode

import (
	"container/heap"
	"sort"
)

func maximumProduct(nums []int, k int) int {
	ans, h := 1, intHeap{nums}
	for heap.Init(&h); k > 0; k-- {
		h.IntSlice[0]++ // add 1 to the minimum
		heap.Fix(&h, 0)
	}
	for _, num := range nums {
		ans = ans * num % (1e9 + 7)
	}
	return ans
}

type intHeap struct{ sort.IntSlice }

func (intHeap) Push(interface{})     {}
func (intHeap) Pop() (_ interface{}) { return }
