package _451_Sort_Characters_By_Frequency

import (
	"container/heap"
	"sort"
)

type kv struct {
	b     rune
	count int
}

func frequencySort(s string) string {
	ans, rune2cnt := make([]rune, 0, len(s)), map[rune]int{}
	for _, c := range s {
		rune2cnt[c]++
	}
	kvSlice := make([]kv, 0, len(rune2cnt))
	for k, v := range rune2cnt {
		kvSlice = append(kvSlice, kv{k, v})
	}
	sort.Slice(kvSlice, func(i, j int) bool { return kvSlice[i].count > kvSlice[j].count })

	for _, kvs := range kvSlice {
		for i := 0; i < kvs.count; i++ {
			ans = append(ans, kvs.b)
		}
	}
	return string(ans)
}

func frequencySort2(s string) string {
	ans, rune2cnt := make([]rune, 0, len(s)), map[rune]int{}
	for _, c := range s {
		rune2cnt[c]++
	}
	h := make(IntHeap, 0, len(rune2cnt))
	for ascii, cnt := range rune2cnt {
		h = append(h, kv{ascii, cnt})
	}
	heap.Init(&h)
	for h.Len() > 0 {
		kvs := heap.Pop(&h).(kv)
		for i := 0; i < kvs.count; i++ {
			ans = append(ans, kvs.b)
		}
	}
	return string(ans)
}

type IntHeap []kv

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
