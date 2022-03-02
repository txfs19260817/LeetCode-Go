package _675_Minimize_Deviation_in_Array

import "container/heap"

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *intHeap) Top() interface{} {
	return (*h)[0]
}

func minimumDeviation(nums []int) int {
	largeH, mini := &intHeap{}, 1<<31
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		heap.Push(largeH, num)
		mini = min(mini, num)
	}
	ans := largeH.Top().(int) - mini
	for largeH.Len() > 0 && largeH.Top().(int)%2 == 0 {
		top := heap.Pop(largeH).(int) / 2
		heap.Push(largeH, top)
		mini = min(mini, top)
		ans = min(ans, largeH.Top().(int)-mini)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
