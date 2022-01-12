package _347_Top_K_Frequent_Elements

import "container/heap"

type pair struct {
	num, cnt int
}

type pairMinHeap []pair

func (h pairMinHeap) Len() int {
	return len(h)
}

func (h pairMinHeap) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt
}

func (h pairMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairMinHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func (h *pairMinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func topKFrequent(nums []int, k int) []int {
	var ans []int
	num2cnt, pHeap := map[int]int{}, make(pairMinHeap, 0, k+1)
	for _, num := range nums {
		num2cnt[num]++
	}
	for num, cnt := range num2cnt {
		heap.Push(&pHeap, pair{num, cnt})
		if pHeap.Len() > k {
			heap.Pop(&pHeap)
		}
	}
	for _, p := range pHeap {
		ans = append(ans, p.num)
	}
	return ans
}
