package snowflake

type Trie struct {
	children [26]*Trie
	value    map[string]int
}

func (t *Trie) Insert(key string, value int) {
	p := t
	for _, ch := range key {
		ch -= 'a'
		if p.children[ch] == nil {
			p.children[ch] = &Trie{}
		}
		p = p.children[ch]
		if p.value == nil {
			p.value = make(map[string]int)
		}
		p.value[key] = value
	}
}

func (t *Trie) Search(key string) map[string]int {
	p := t
	for _, ch := range key {
		ch -= 'a'
		if p.children[ch] == nil {
			return nil
		}
		p = p.children[ch]
	}
	return p.value
}

func (t *Trie) Delete(key string) {
	p := t
	for _, ch := range key {
		ch -= 'a'
		if p.children[ch] == nil {
			return
		}
		p = p.children[ch]
		delete(p.value, key)
	}
}

type KVStore struct {
	trie *Trie
}

func NewKVStore() KVStore {
	return KVStore{trie: &Trie{}}
}

func (s *KVStore) Set(key string, value int) {
	s.trie.Insert(key, value)
}

func (s *KVStore) Get(key string) int {
	res := s.trie.Search(key)
	if val, ok := res[key]; ok {
		return val
	}
	return -1
}

func (s *KVStore) Update(key string, value int) {
	s.Set(key, value)
}

func (s *KVStore) DeleteKey(key string) {
	s.trie.Delete(key)
}

func (s *KVStore) PrefixSearch(prefix string) []int {
	res := s.trie.Search(prefix)
	ans := make([]int, 0, len(res))
	for _, v := range res {
		ans = append(ans, v)
	}
	return ans
}
