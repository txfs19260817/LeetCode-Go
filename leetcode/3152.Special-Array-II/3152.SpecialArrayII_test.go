package leetcode

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example 1",
			args: args{
				nums:    []int{3, 4, 1, 2, 6},
				queries: [][]int{{0, 4}},
			},
			want: []bool{false},
		},
		{
			name: "Example 2",
			args: args{
				nums:    []int{4, 3, 1, 6},
				queries: [][]int{{0, 2}, {2, 3}},
			},
			want: []bool{false, true},
		},
		{
			name: "Alternating parity stays special across all ranges",
			args: args{
				nums:    []int{1, 2, 3, 4, 5, 6},
				queries: [][]int{{0, 5}, {1, 4}, {2, 2}},
			},
			want: []bool{true, true, true},
		},
		{
			name: "Same parity adjacency breaks overlapping ranges",
			args: args{
				nums:    []int{2, 4, 5, 7, 8},
				queries: [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}},
			},
			want: []bool{false, false, false, true},
		},
		{
			name: "Single element query is always special",
			args: args{
				nums:    []int{8, 10, 3},
				queries: [][]int{{0, 0}, {1, 1}, {2, 2}},
			},
			want: []bool{true, true, true},
		},
		{
			name: "Bad pair at the end only affects suffix ranges",
			args: args{
				nums:    []int{1, 2, 5, 8, 10},
				queries: [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 4}},
			},
			want: []bool{true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
