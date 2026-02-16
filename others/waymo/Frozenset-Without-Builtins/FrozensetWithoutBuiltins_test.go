package waymo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrozenSetExampleFlow(t *testing.T) {
	s1 := NewFrozenSet([]int{5, 1, 3, 3})
	s2 := s1.With(4)
	s3 := s2.Without(1)
	s4 := NewFrozenSet([]int{3, 7})

	assert.Equal(t, []int{1, 3, 5}, s1.Elements())
	assert.Equal(t, []int{1, 3, 4, 5}, s2.Elements())
	assert.Equal(t, []int{3, 4, 5}, s3.Elements())
	assert.Equal(t, []int{3, 4, 5, 7}, s3.Union(s4).Elements())
	assert.Equal(t, []int{3}, s3.Intersection(s4).Elements())
	assert.Equal(t, []int{4, 5}, s3.Difference(s4).Elements())
}

func TestFrozenSetTable(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		query  int
		size   int
		want   []int
	}{
		{
			name:   "empty",
			values: []int{},
			query:  1,
			size:   0,
			want:   []int{},
		},
		{
			name:   "deduplicated and sorted",
			values: []int{9, 3, 5, 3, 1, 9},
			query:  5,
			size:   4,
			want:   []int{1, 3, 5, 9},
		},
		{
			name:   "negative values",
			values: []int{0, -3, -1, 2, -3},
			query:  -1,
			size:   4,
			want:   []int{-3, -1, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewFrozenSet(tt.values)
			assert.Equal(t, tt.size, set.Size())
			assert.Equal(t, tt.want, set.Elements())
			assert.Equal(t, containsInt(tt.want, tt.query), set.Contains(tt.query))
		})
	}
}

func TestFrozenSetImmutability(t *testing.T) {
	original := NewFrozenSet([]int{2, 4, 6})
	afterAdd := original.With(8)
	afterRemove := afterAdd.Without(4)

	assert.Equal(t, []int{2, 4, 6}, original.Elements())
	assert.Equal(t, []int{2, 4, 6, 8}, afterAdd.Elements())
	assert.Equal(t, []int{2, 6, 8}, afterRemove.Elements())
	assert.False(t, original.Contains(8))
	assert.True(t, afterAdd.Contains(8))
	assert.False(t, afterRemove.Contains(4))
}

func TestFrozenSetSetOpsAndEquals(t *testing.T) {
	left := NewFrozenSet([]int{1, 2, 3, 5})
	right := NewFrozenSet([]int{3, 4, 5, 6})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, left.Union(right).Elements())
	assert.Equal(t, []int{3, 5}, left.Intersection(right).Elements())
	assert.Equal(t, []int{1, 2}, left.Difference(right).Elements())
	assert.True(t, left.Equals(NewFrozenSet([]int{5, 3, 2, 1})))
	assert.False(t, left.Equals(right))
}

func TestFrozenSetLargeSortedInsertions(t *testing.T) {
	set := NewFrozenSet([]int{})
	for i := 1; i <= 300; i++ {
		set = set.With(i)
	}
	assert.Equal(t, 300, set.Size())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, set.Elements()[:5])
	assert.Equal(t, []int{296, 297, 298, 299, 300}, set.Elements()[295:])
}

func containsInt(values []int, target int) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
