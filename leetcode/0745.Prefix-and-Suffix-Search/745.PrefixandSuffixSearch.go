package _745_Prefix_and_Suffix_Search

import "strings"

type WordFilter struct {
	words []string
	w2idx map[string]int
	trie  *Trie
}

func Constructor(words []string) WordFilter {
	w2idx, trie := map[string]int{}, NewTrie()
	for i, word := range words {
		w2idx[word] = i
		trie.insert(word)
	}
	return WordFilter{words, w2idx, trie}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	ans, t := -1, this.trie.searchPrefix(prefix)
	if t == nil {
		return ans
	}
	var dfs func(t *Trie, path []byte)
	dfs = func(t *Trie, path []byte) {
		if t == nil {
			return
		}
		if t.isEnd {
			if fullWord := prefix + string(path); strings.HasSuffix(fullWord, suffix) && ans < this.w2idx[fullWord] {
				ans = this.w2idx[fullWord]
			}
			return
		}
		for i, child := range t.children {
			if child != nil {
				path = append(path, byte(i)+'a')
				dfs(child, path)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(t, []byte{})
	return ans
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) insert(word string) {
	p := t
	for _, ch := range word {
		ch -= 'a'
		if p.children[ch] == nil {
			p.children[ch] = &Trie{}
		}
		p = p.children[ch]
	}
	p.isEnd = true
}

func (t *Trie) searchPrefix(word string) *Trie {
	p := t
	for _, ch := range word {
		ch -= 'a'
		if p.children[ch] == nil {
			return nil
		}
		p = p.children[ch]
	}
	return p
}
