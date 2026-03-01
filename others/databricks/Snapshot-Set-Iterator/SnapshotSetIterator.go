package databricks

const (
	keyEntry = 0
	keyIndex = 1
)

type storeKey struct {
	kind int
	id   int
}

type storeVal struct {
	a int
	b int
}

// SnapshotSet is a set of integers that supports snapshotted iteration.
type SnapshotSet struct {
	// Single map storage:
	// {keyEntry, addVersion} -> {value, removeVersion}
	// {keyIndex, value} -> {latestAddVersion, _}
	store   map[storeKey]storeVal
	version int
}

// NewSnapshotSet creates a new empty SnapshotSet.
func NewSnapshotSet() *SnapshotSet {
	return &SnapshotSet{
		store: make(map[storeKey]storeVal),
	}
}

// Add inserts n into the set. Returns true if newly added, false if already present.
func (s *SnapshotSet) Add(n int) bool {
	idxKey := storeKey{kind: keyIndex, id: n}
	if idxVal, ok := s.store[idxKey]; ok {
		entKey := storeKey{kind: keyEntry, id: idxVal.a}
		if entVal, ok := s.store[entKey]; ok && entVal.b == -1 {
			return false // already present and not removed
		}
	}

	addVersion := s.version
	s.store[storeKey{kind: keyEntry, id: addVersion}] = storeVal{a: n, b: -1}
	s.store[idxKey] = storeVal{a: addVersion}
	s.version++
	return true
}

// Remove removes n from the set. Returns true if removed, false if not found.
func (s *SnapshotSet) Remove(n int) bool {
	idxVal, ok := s.store[storeKey{kind: keyIndex, id: n}]
	if !ok {
		return false
	}
	entKey := storeKey{kind: keyEntry, id: idxVal.a}
	entVal, ok := s.store[entKey]
	if !ok || entVal.b != -1 {
		return false
	}

	entVal.b = s.version
	s.store[entKey] = entVal
	s.version++
	return true
}

// Contains returns true if n is currently in the set.
func (s *SnapshotSet) Contains(n int) bool {
	idxVal, ok := s.store[storeKey{kind: keyIndex, id: n}]
	if !ok {
		return false
	}
	entVal, ok := s.store[storeKey{kind: keyEntry, id: idxVal.a}]
	if !ok {
		return false
	}
	return entVal.b == -1
}

// GetIterator returns an iterator over the elements present at this moment,
// in insertion order. Later mutations do not affect this iterator.
func (s *SnapshotSet) GetIterator() *SnapshotIterator {
	return &SnapshotIterator{
		set:         s,
		snapVersion: s.version,
		scanLimit:   s.version,
	}
}

// SnapshotIterator iterates over a frozen snapshot of the set.
// It performs lazy scanning and uses O(1) iterator space.
type SnapshotIterator struct {
	set         *SnapshotSet
	snapVersion int
	scanLimit   int
	scanPos     int
	buffered    bool
	buffer      int
}

// HasNext returns true if there are more elements to iterate.
func (it *SnapshotIterator) HasNext() bool {
	if it.buffered {
		return true
	}

	for it.scanPos < it.scanLimit {
		addVersion := it.scanPos
		it.scanPos++
		entVal, ok := it.set.store[storeKey{kind: keyEntry, id: addVersion}]
		if !ok {
			continue
		}
		if entVal.b == -1 || entVal.b >= it.snapVersion {
			it.buffer = entVal.a
			it.buffered = true
			return true
		}
	}
	return false
}

// Next returns the next element. Panics if no next element exists.
func (it *SnapshotIterator) Next() int {
	if !it.HasNext() {
		panic("SnapshotIterator: no next element")
	}
	val := it.buffer
	it.buffered = false
	return val
}
