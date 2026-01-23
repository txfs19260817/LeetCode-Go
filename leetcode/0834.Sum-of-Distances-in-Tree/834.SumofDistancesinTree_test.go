package leetcode

import (
	"reflect"
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			name: "example2_single_node",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "example3_two_nodes",
			args: args{
				n:     2,
				edges: [][]int{{1, 0}},
			},
			want: []int{1, 1},
		},
		{
			name: "three_node_chain",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {1, 2}},
			},
			want: []int{3, 2, 3},
		},
		{
			name: "four_node_star",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}},
			},
			want: []int{3, 5, 5, 5},
		},
		{
			name: "unbalanced_tree",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			},
			want: []int{8, 5, 8, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDistancesInTree(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
