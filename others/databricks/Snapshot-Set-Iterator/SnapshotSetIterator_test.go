package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// collectIterator drains a SnapshotIterator into a slice.
func collectIterator(it *SnapshotIterator) []int {
	result := []int{}
	for it.HasNext() {
		result = append(result, it.Next())
	}
	return result
}

func TestSnapshotSet_MainExample(t *testing.T) {
	set := NewSnapshotSet()

	assert.True(t, set.Add(1))
	assert.True(t, set.Add(2))
	assert.True(t, set.Add(3))
	assert.True(t, set.Add(4))
	assert.False(t, set.Add(1)) // already exists

	it1 := set.GetIterator() // snapshot: [1,2,3,4]

	assert.True(t, set.Remove(1))
	assert.True(t, set.Remove(3))
	assert.False(t, set.Remove(5)) // not present

	it2 := set.GetIterator() // snapshot: [2,4]

	assert.Equal(t, []int{1, 2, 3, 4}, collectIterator(it1))
	assert.Equal(t, []int{2, 4}, collectIterator(it2))
}

func TestSnapshotSet_ReAddAfterRemove(t *testing.T) {
	set := NewSnapshotSet()

	assert.True(t, set.Add(1))
	assert.True(t, set.Add(2))
	assert.True(t, set.Remove(1))
	assert.False(t, set.Contains(1))

	// Re-add 1 — it should appear after 2 in insertion order
	assert.True(t, set.Add(1))
	assert.True(t, set.Contains(1))

	it := set.GetIterator()
	assert.Equal(t, []int{2, 1}, collectIterator(it))
}

func TestSnapshotSet_MultipleConcurrentIterators(t *testing.T) {
	set := NewSnapshotSet()

	assert.True(t, set.Add(10))
	assert.True(t, set.Add(20))

	it1 := set.GetIterator() // snapshot: [10, 20]

	assert.True(t, set.Add(30))

	it2 := set.GetIterator() // snapshot: [10, 20, 30]

	assert.True(t, set.Remove(10))

	it3 := set.GetIterator() // snapshot: [20, 30]

	// All three iterators work independently
	assert.Equal(t, []int{10, 20}, collectIterator(it1))
	assert.Equal(t, []int{10, 20, 30}, collectIterator(it2))
	assert.Equal(t, []int{20, 30}, collectIterator(it3))
}

func TestSnapshotSet_EmptySet(t *testing.T) {
	set := NewSnapshotSet()

	it := set.GetIterator()
	assert.False(t, it.HasNext())
	assert.Equal(t, []int{}, collectIterator(it))
}

func TestSnapshotSet_NextPanicsWhenEmpty(t *testing.T) {
	set := NewSnapshotSet()
	it := set.GetIterator()

	assert.Panics(t, func() { it.Next() })
}
