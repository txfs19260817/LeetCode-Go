package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	nestedList := []*NestedInteger{
		makeList(
			makeInt(1),
			makeInt(1),
		),
		makeInt(2),
		makeList(
			makeInt(1),
			makeInt(1),
		),
	}

	assert.Equal(t, []int{1, 1, 2, 1, 1}, collectIterator(Constructor(nestedList)))
	assert.Equal(t, []int{1, 1, 2, 1, 1}, collectIterator2(Constructor2(nestedList)))
}

func TestConstructorEmpty(t *testing.T) {
	assert.False(t, Constructor(nil).HasNext())
	assert.False(t, Constructor2(nil).HasNext())
}

func TestConstructorDeepNested(t *testing.T) {
	nestedList := []*NestedInteger{
		makeInt(1),
		makeList(
			makeInt(4),
			makeList(
				makeInt(6),
			),
		),
	}

	assert.Equal(t, []int{1, 4, 6}, collectIterator(Constructor(nestedList)))
	assert.Equal(t, []int{1, 4, 6}, collectIterator2(Constructor2(nestedList)))
}

func BenchmarkConstructor(b *testing.B) {
	nestedList := makeBenchmarkNestedList(256)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = collectIterator(Constructor(nestedList))
	}
}

func BenchmarkConstructor2(b *testing.B) {
	nestedList := makeBenchmarkNestedList(256)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = collectIterator2(Constructor2(nestedList))
	}
}

func makeInt(v int) *NestedInteger {
	n := &NestedInteger{}
	n.SetInteger(v)
	return n
}

func makeList(items ...*NestedInteger) *NestedInteger {
	n := &NestedInteger{}
	for _, item := range items {
		n.Add(*item)
	}
	return n
}

func collectIterator(it *NestedIterator) []int {
	var ans []int
	for it.HasNext() {
		ans = append(ans, it.Next())
	}
	return ans
}

func collectIterator2(it *NestedIterator2) []int {
	var ans []int
	for it.HasNext() {
		ans = append(ans, it.Next())
	}
	return ans
}

func makeBenchmarkNestedList(size int) []*NestedInteger {
	nestedList := make([]*NestedInteger, 0, size)
	for i := 0; i < size; i += 4 {
		nestedList = append(nestedList, makeList(
			makeInt(i),
			makeList(
				makeInt(i+1),
				makeList(
					makeInt(i+2),
				),
			),
			makeInt(i+3),
		))
	}
	return nestedList
}
