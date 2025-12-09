package leetcode

import "strconv"

type Trie struct {
	children [26]*Trie
	count    int
}

func (t *Trie) Insert(word string) {
	p := t
	for _, r := range word {
		idx := r - 'a'
		if p.children[idx] == nil {
			p.children[idx] = &Trie{}
		}
		p.count++
		p = p.children[idx]
	}
}

func (t *Trie) PrefixLen(word string) (cnt int) {
	p := t
	for _, r := range word {
		idx := r - 'a'
		if p.children[idx].count == 1 {
			break
		}
		cnt++
		p = p.children[idx]
	}
	return
}

type WordIdx struct {
	word string
	idx  int
}

func wordsAbbreviation(words []string) []string {
	ans := make([]string, len(words))
	abbv2words := map[string][]WordIdx{}
	for i, word := range words {
		a := getAbbreviation(word, 0)
		abbv2words[a] = append(abbv2words[a], WordIdx{word, i})
	}
	for _, group := range abbv2words {
		t := &Trie{}

		// insert
		for _, wi := range group {
			t.Insert(wi.word)
		}
		// retrieve
		for _, wi := range group {
			cnt := t.PrefixLen(wi.word)
			ans[wi.idx] = getAbbreviation(wi.word, cnt)
		}
	}
	return ans
}

func getAbbreviation(word string, keep int) string {
	n := len(word)
	if n-keep-2 <= 1 { // never let strconv.Itoa(n-keep-2) <= "1"
		return word
	}
	return word[:keep+1] + strconv.Itoa(n-keep-2) + string(word[n-1])
}
