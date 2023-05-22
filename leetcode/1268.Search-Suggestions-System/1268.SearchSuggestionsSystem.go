package leetcode

import "container/heap"

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
