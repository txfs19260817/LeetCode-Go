package leetcode

import (
	"bytes"
	"container/heap"
)

type pair struct {
	b byte
	n int
}

type pairHeap []pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].b > h[j].b
}

func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *pairHeap) Top() interface{} {
	return (*h)[0]
}

func repeatLimitedString(s string, repeatLimit int) string {
	ans, b2n, h := make([]byte, 0, len(s)), map[byte]int{}, &pairHeap{}
	for i := 0; i < len(s); i++ {
		b2n[s[i]]++
	}
	for b, n := range b2n {
		heap.Push(h, pair{b, n})
	}
	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		if top.n <= repeatLimit {
			ans = append(ans, bytes.Repeat([]byte{top.b}, top.n)...)
			continue
		}
		if h.Len() == 0 {
			ans = append(ans, bytes.Repeat([]byte{top.b}, min(repeatLimit, top.n))...)
			break
		}
		ans = append(ans, bytes.Repeat([]byte{top.b}, repeatLimit)...)
		top.n -= repeatLimit
		sub := heap.Pop(h).(pair)
		ans = append(ans, sub.b)
		sub.n--
		if sub.n > 0 {
			heap.Push(h, sub)
		}
		heap.Push(h, top)
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
