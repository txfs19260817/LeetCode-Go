package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "k = 3, n = 7",
			args: args{3, 7},
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "k = 3, n = 9",
			args: args{3, 9},
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "k = 4, n = 1",
			args: args{4, 1},
			want: nil,
		},
		{
			name: "k = 3, n = 2",
			args: args{3, 2},
			want: nil,
		},
		{
			name: "k = 9, n = 45",
			args: args{9, 45},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect := normalizeCombinations(tt.want)
			actual := normalizeCombinations(combinationSum3(tt.args.k, tt.args.n))
			assert.Equal(t, expect, actual)
			actual = normalizeCombinations(combinationSum32(tt.args.k, tt.args.n))
			assert.Equal(t, expect, actual)
		})
	}
}

func Benchmark_combinationSum3(b *testing.B) {
	cases := combinationSum3BenchmarkCases()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		for _, tc := range cases {
			combinationSum3(tc.k, tc.n)
		}
	}
}

func Benchmark_combinationSum32(b *testing.B) {
	cases := combinationSum3BenchmarkCases()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		for _, tc := range cases {
			combinationSum32(tc.k, tc.n)
		}
	}
}

func combinationSum3BenchmarkCases() []struct {
	k int
	n int
} {
	return []struct {
		k int
		n int
	}{
		{k: 3, n: 7},
		{k: 3, n: 9},
		{k: 4, n: 18},
		{k: 5, n: 21},
		{k: 6, n: 30},
		{k: 7, n: 35},
		{k: 8, n: 36},
		{k: 9, n: 45},
	}
}

func normalizeCombinations(nums [][]int) [][]int {
	if nums == nil {
		return nil
	}

	normalized := make([][]int, len(nums))
	for i, combination := range nums {
		normalized[i] = slices.Clone(combination)
		slices.Sort(normalized[i])
	}

	slices.SortFunc(normalized, func(a, b []int) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return len(a) - len(b)
	})

	return normalized
}
