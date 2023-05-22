package leetcode

import "strings"

type Encrypter struct {
	dict Trie
	k2v  map[byte]string
	v2ks map[string][]byte
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	e := Encrypter{
		dict: NewTrie(),
		k2v:  make(map[byte]string),
		v2ks: make(map[string][]byte),
	}
	for i, k := range keys {
		v := values[i]
		e.v2ks[v] = append(e.v2ks[v], k)
		e.k2v[k] = v
	}
	for _, d := range dictionary {
		e.dict.Insert(d)
	}
	return e
}

func (this *Encrypter) Encrypt(word1 string) string {
	ans := make([]string, 0, len(word1))
	for _, r := range word1 {
		ans = append(ans, this.k2v[byte(r)])
	}
	return strings.Join(ans, "")
}

func (this *Encrypter) Decrypt(word2 string) int {
	var cnt int
	this.dfs(word2, 0, []byte{}, &this.dict, &cnt)
	return cnt
}

func (this *Encrypter) dfs(word2 string, i int, path []byte, trie *Trie, cnt *int) {
	if i == len(word2) {
		if this.dict.Search(string(path)) {
			*cnt++
		}
		return
	}
	v := word2[i : i+2]
	keys := this.v2ks[v]
	for _, k := range keys {
		if nextT := trie.SearchPrefix(string(k)); nextT != nil {
			path = append(path, k)
			this.dfs(word2, i+2, path, nextT, cnt)
			path = path[:len(path)-1]
		}
	}
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// NewTrie initializes a Trie.
func NewTrie() Trie {
	return Trie{}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
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

func (t *Trie) SearchPrefix(word string) *Trie {
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

// Search returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd == true
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
