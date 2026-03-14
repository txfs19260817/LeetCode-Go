package leetcode

import (
	"container/heap"
	"sort"
)

type strHeap []string

func (h strHeap) Len() int {
	return len(h)
}

func (h strHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h strHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *strHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *strHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

type trie struct {
	children    [26]*trie
	suggestions *strHeap
}

func newTrie() *trie {
	return &trie{suggestions: &strHeap{}}
}

func (t *trie) Insert(word string) {
	p := t
	for _, c := range word {
		c -= 'a'
		if p.children[c] == nil {
			p.children[c] = newTrie()
		}
		heap.Push(p.children[c].suggestions, word)
		if p.children[c].suggestions.Len() > 3 {
			heap.Pop(p.children[c].suggestions)
		}
		p = p.children[c]
	}
}

func (t *trie) GetSuggestions(word string) [][]string {
	ans, p := make([][]string, len(word)), t
	for i, c := range word {
		c -= 'a'
		if p.children[c] == nil {
			break
		}
		sug := make([]string, 0, 3)
		for i := 0; i < 3 && p.children[c].suggestions.Len() > 0; i++ {
			sug = append(sug, heap.Pop(p.children[c].suggestions).(string))
		}
		if len(sug) > 1 {
			sug[0], sug[len(sug)-1] = sug[len(sug)-1], sug[0]
		}
		ans[i] = sug
		p = p.children[c]
	}
	return ans
}

func suggestedProducts(products []string, searchWord string) [][]string {
	t := newTrie()
	for _, w := range products {
		t.Insert(w)
	}
	return t.GetSuggestions(searchWord)
}

func suggestedProducts2(products []string, searchWord string) [][]string {
	sort.Strings(products)
	ans := make([][]string, 0, len(searchWord))
	for i := range searchWord {
		query := searchWord[:i+1]
		j := sort.Search(len(products), func(i int) bool { return query <= products[i] })
		cur := make([]string, 0, 3)
		for ; j < len(products) && len(query) <= len(products[j]) && products[j][:i+1] == query; j++ {
			cur = append(cur, products[j])
			if len(cur) >= 3 {
				break
			}
		}
		ans = append(ans, cur)
	}
	return ans
}
