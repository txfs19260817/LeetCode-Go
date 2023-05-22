package leetcode

type MapSum struct {
	t *trie
}

func Constructor() MapSum {
	return MapSum{newTrie()}
}

func (this *MapSum) Insert(key string, val int) {
	this.t.insert(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	return this.t.find(prefix)
}

type trie struct {
	children [26]*trie
	dict     map[string]int
}

func newTrie() *trie {
	return &trie{dict: make(map[string]int)}
}

func (t *trie) insert(key string, val int) {
	p := t
	for _, ch := range key {
		ch -= 'a'
		if p.children[ch] == nil {
			p.children[ch] = newTrie()
		}
		p = p.children[ch]
		p.dict[key] = val
	}
}

func (t *trie) find(key string) (val int) {
	p := t
	for _, ch := range key {
		ch -= 'a'
		if p.children[ch] == nil {
			return 0
		}
		p = p.children[ch]
	}
	for _, v := range p.dict {
		val += v
	}
	return
}
