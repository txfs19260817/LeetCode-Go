package databricks

type entry struct {
	value         int
	addVersion    int
	removeVersion int // -1 means not removed
}

// SnapshotSet is a set of integers that supports snapshotted iteration.
type SnapshotSet struct {
	entries []entry
	index   map[int]int // value -> index in entries
	version int
}

// NewSnapshotSet creates a new empty SnapshotSet.
func NewSnapshotSet() *SnapshotSet {
	return &SnapshotSet{
		index: make(map[int]int),
	}
}

// Add inserts n into the set. Returns true if newly added, false if already present.
func (s *SnapshotSet) Add(n int) bool {
	if idx, ok := s.index[n]; ok && s.entries[idx].removeVersion == -1 {
		return false // already present and not removed
	}
	s.entries = append(s.entries, entry{
		value:         n,
		addVersion:    s.version,
		removeVersion: -1,
	})
	s.index[n] = len(s.entries) - 1
	s.version++
	return true
}

// Remove removes n from the set. Returns true if removed, false if not found.
func (s *SnapshotSet) Remove(n int) bool {
	idx, ok := s.index[n]
	if !ok || s.entries[idx].removeVersion != -1 {
		return false
	}
	s.entries[idx].removeVersion = s.version
	s.version++
	return true
}

// Contains returns true if n is currently in the set.
func (s *SnapshotSet) Contains(n int) bool {
	idx, ok := s.index[n]
	if !ok {
		return false
	}
	return s.entries[idx].removeVersion == -1
}

// GetIterator returns an iterator over the elements present at this moment,
// in insertion order. Later mutations do not affect this iterator.
func (s *SnapshotSet) GetIterator() *SnapshotIterator {
	snapVersion := s.version
	snapshot := []int{}
	for _, e := range s.entries {
		if e.addVersion < snapVersion && (e.removeVersion == -1 || e.removeVersion >= snapVersion) {
			snapshot = append(snapshot, e.value)
		}
	}
	return &SnapshotIterator{elements: snapshot}
}

// SnapshotIterator iterates over a frozen snapshot of the set.
type SnapshotIterator struct {
	elements []int
	pos      int
}

// HasNext returns true if there are more elements to iterate.
func (it *SnapshotIterator) HasNext() bool {
	return it.pos < len(it.elements)
}

// Next returns the next element. Panics if no next element exists.
func (it *SnapshotIterator) Next() int {
	if !it.HasNext() {
		panic("SnapshotIterator: no next element")
	}
	val := it.elements[it.pos]
	it.pos++
	return val
}
