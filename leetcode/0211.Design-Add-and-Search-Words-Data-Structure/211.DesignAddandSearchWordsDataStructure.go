package leetcode

import "strings"

type WordDictionary struct {
	triePre, trieSuf *Trie
	hashTable        map[string]bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		triePre:   NewTrie(),
		trieSuf:   NewTrie(),
		hashTable: make(map[string]bool),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.hashTable[word] = true
	this.triePre.Add(word)
	this.trieSuf.Add(reverseString(word))
}

func (this *WordDictionary) Search(word string) bool {
	// no dot
	if !strings.ContainsRune(word, '.') {
		return this.hashTable[word]
	}

	// all dots
	trimLeft := strings.TrimLeft(word, ".")
	if len(trimLeft) == 0 {
		for k := range this.hashTable {
			if len(word) == len(k) {
				return true
			}
		}
		return false
	}

	if word[0] == '.' { // prefix dots
		if !strings.ContainsRune(trimLeft, '.') {
			return this.trieSuf.SearchPrefix(reverseString(trimLeft), len(word))
		}
	} else if word[len(word)-1] == '.' { // suffix dots
		trimRight := strings.TrimRight(word, ".")
		if !strings.ContainsRune(trimRight, '.') {
			return this.triePre.SearchPrefix(trimRight, len(word))
		}
	}

	// others: scan
	for k := range this.hashTable {
		if len(word) != len(k) {
			continue
		}
		found := true
		for i := 0; i < len(word); i++ {
			if word[i] != '.' && word[i] != k[i] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

type Trie struct {
	children    [26]*Trie
	wordLengths [26]map[int]bool
}

func NewTrie() *Trie {
	var m [26]map[int]bool
	for i := 0; i < 26; i++ {
		m[i] = make(map[int]bool)
	}
	return &Trie{wordLengths: m}
}

func (t *Trie) Add(word string) {
	p := t
	for _, ch := range word {
		ch -= 'a'
		if p.children[ch] == nil {
			p.children[ch] = NewTrie()
		}
		p.wordLengths[ch][len(word)] = true
		p = p.children[ch]
	}
}

func (t *Trie) SearchPrefix(word string, expectedLen int) bool {
	p, res := t, false
	for _, ch := range word {
		ch -= 'a'
		if p.children[ch] == nil {
			return false
		}
		res = p.wordLengths[ch][expectedLen]
		p = p.children[ch]
	}
	return res
}

func reverseString(str string) string {
	s := []byte(str)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	return string(s)
}
