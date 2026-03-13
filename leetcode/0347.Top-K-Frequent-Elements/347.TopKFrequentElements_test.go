package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		wants [][]int
	}{
		{
			name:  "nums = [1,1,1,2,2,3], k = 2",
			args:  args{[]int{1, 1, 1, 2, 2, 3}, 2},
			wants: [][]int{{1, 2}},
		},
		{
			name:  "nums = [1], k = 1",
			args:  args{[]int{1}, 1},
			wants: [][]int{{1}},
		},
		{
			name:  "nums = [5,5,5,5], k = 1",
			args:  args{[]int{5, 5, 5, 5}, 1},
			wants: [][]int{{5}},
		},
		{
			name:  "nums = [4,1,-1,2,-1,2,3], k = 2",
			args:  args{[]int{4, 1, -1, 2, -1, 2, 3}, 2},
			wants: [][]int{{-1, 2}},
		},
		{
			name:  "nums = [1,1,2,2,3], k = 3",
			args:  args{[]int{1, 1, 2, 2, 3}, 3},
			wants: [][]int{{1, 2, 3}},
		},
		{
			name:  "nums = [1,1,2,2,3,3], k = 2",
			args:  args{[]int{1, 1, 2, 2, 3, 3}, 2},
			wants: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertAnyElementsMatch(t, tt.wants, topKFrequent(tt.args.nums, tt.args.k))
			assertAnyElementsMatch(t, tt.wants, topKFrequent2(tt.args.nums, tt.args.k))
		})
	}
}

func Benchmark_topKFrequent(b *testing.B) {
	nums := benchmarkTopKFrequentInput()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent(nums, 128)
	}
}

func Benchmark_topKFrequent2(b *testing.B) {
	nums := benchmarkTopKFrequentInput()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent2(nums, 128)
	}
}

func benchmarkTopKFrequentInput() []int {
	const uniqueValues = 50000
	const maxFrequency = 32

	nums := make([]int, 0, uniqueValues*(maxFrequency+1)/2)
	for value := 1; value <= uniqueValues; value++ {
		freq := 1 + value%maxFrequency
		for repeat := 0; repeat < freq; repeat++ {
			nums = append(nums, value)
		}
	}
	return nums
}

func assertAnyElementsMatch(t *testing.T, wants [][]int, got []int) {
	t.Helper()

	for _, want := range wants {
		if sameElements(want, got) {
			return
		}
	}

	assert.Failf(t, "unexpected result", "got %v, want one of %v", got, wants)
}

func sameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[int]int, len(a))
	for _, n := range a {
		counts[n]++
	}
	for _, n := range b {
		counts[n]--
		if counts[n] < 0 {
			return false
		}
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
