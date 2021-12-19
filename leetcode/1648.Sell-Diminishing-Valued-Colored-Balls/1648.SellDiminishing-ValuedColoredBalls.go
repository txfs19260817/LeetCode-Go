package _648_Sell_Diminishing_Valued_Colored_Balls

import "container/heap"

const mod = 1000000007

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
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func (h *intHeap) Top() interface{} {
	return (*h)[0]
}

func maxProfit(inventory []int, orders int) int {
	ans, h := 0, intHeap(inventory)
	heap.Init(&h)
	for orders > 0 {
		top, topCnt := heap.Pop(&h).(int), 1
		for h.Len() > 0 && top == h.Top().(int) {
			topCnt++
			heap.Pop(&h)
		}
		var delta, nextTop int
		if h.Len() > 0 {
			nextTop = h.Top().(int)
			delta = topCnt * (top - nextTop) // e.g. [3,3,2] => ans += 3+3 => heap = [2,2,2]
		}
		if h.Len() == 0 || orders < delta { // h.Len() == 0 => all elements are equal || orders < delta => needed inventory = min(orders, delta)
			offset, extra := orders/topCnt, orders%topCnt
			subTop := top - offset
			ans = (ans + topCnt*(top+subTop+1)*offset/2) % mod
			ans = (ans + extra*subTop) % mod
			break
		}
		ans = (ans + (top+nextTop+1)*delta/2) % mod
		orders -= delta
		for ; topCnt > 0; topCnt-- {
			heap.Push(&h, nextTop)
		}
	}
	return ans
}
