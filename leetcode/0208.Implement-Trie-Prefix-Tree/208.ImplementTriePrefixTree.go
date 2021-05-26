package _208_Implement_Trie_Prefix_Tree

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// Constructor initializee your data structure here.
func Constructor() Trie {
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

// Search returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd == true
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}
