package snowflake

import (
	"strings"
	"sync"
)

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

// KVStore supports thread-safe operations and independent transactions
type KVStore struct {
	mu   sync.RWMutex
	trie *Trie
}

func NewKVStore() *KVStore {
	return &KVStore{trie: &Trie{}}
}

// Transaction represents an isolated atomic operation
type Transaction struct {
	store  *KVStore
	buffer map[string]*int // nil value implies deletion
}

// Begin starts a new independent transaction
func (s *KVStore) Begin() *Transaction {
	return &Transaction{
		store:  s,
		buffer: make(map[string]*int),
	}
}

// Set adds or updates a key in the transaction buffer
func (tx *Transaction) Set(key string, value int) {
	val := value
	tx.buffer[key] = &val
}

// Get retrieves a value, checking the buffer first then the global store
func (tx *Transaction) Get(key string) int {
	// Check local uncommitted changes first
	if val, ok := tx.buffer[key]; ok {
		if val == nil {
			return -1 // Deleted in this transaction
		}
		return *val
	}

	// Read from global store
	tx.store.mu.RLock()
	defer tx.store.mu.RUnlock()

	// We only need exact match for Get, but Trie stores all matches in the path.
	// For exact match, we can just look at Search(key) and check if key exists in the map.
	// Optimization: Trie could have a dedicated Get, but Search(key)[key] works.
	if res := tx.store.trie.Search(key); res != nil {
		if v, ok := res[key]; ok {
			return v
		}
	}
	return -1
}

// DeleteKey marks a key for deletion in the transaction buffer
func (tx *Transaction) DeleteKey(key string) {
	tx.buffer[key] = nil
}

// PrefixSearch returns keys matching prefix, merging global store with local changes
func (tx *Transaction) PrefixSearch(prefix string) []int {
	merged := make(map[string]int)

	// 1. Get snapshot from global store
	tx.store.mu.RLock()
	for k, v := range tx.store.trie.Search(prefix) {
		merged[k] = v
	}
	tx.store.mu.RUnlock()

	// 2. Apply local changes
	for k, v := range tx.buffer {
		if strings.HasPrefix(k, prefix) {
			if v == nil {
				delete(merged, k)
			} else {
				merged[k] = *v
			}
		}
	}

	ans := make([]int, 0, len(merged))
	for _, v := range merged {
		ans = append(ans, v)
	}
	return ans
}

// Commit applies all buffered changes to the global store atomically
func (tx *Transaction) Commit() {
	tx.store.mu.Lock()
	defer tx.store.mu.Unlock()

	for key, val := range tx.buffer {
		if val == nil {
			tx.store.trie.Delete(key)
		} else {
			tx.store.trie.Insert(key, *val)
		}
	}
	// Clear buffer to prevent reuse or duplicate commits
	tx.buffer = make(map[string]*int)
}

// Abort discards the transaction
func (tx *Transaction) Abort() {
	tx.buffer = make(map[string]*int)
}

// --- Direct KVStore methods (Auto-commit) ---

func (s *KVStore) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.trie.Insert(key, value)
}

func (s *KVStore) Get(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if res := s.trie.Search(key); res != nil {
		if v, ok := res[key]; ok {
			return v
		}
	}
	return -1
}

func (s *KVStore) Update(key string, value int) {
	s.Set(key, value)
}

func (s *KVStore) DeleteKey(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.trie.Delete(key)
}

func (s *KVStore) PrefixSearch(prefix string) []int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := s.trie.Search(prefix)
	ans := make([]int, 0, len(res))
	for _, v := range res {
		ans = append(ans, v)
	}
	return ans
}
