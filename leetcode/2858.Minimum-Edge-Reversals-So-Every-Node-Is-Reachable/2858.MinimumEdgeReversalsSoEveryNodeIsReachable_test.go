package leetcode

import (
	"reflect"
	"testing"
)

func Test_minEdgeReversals(t *testing.T) {
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
				n:     4,
				edges: [][]int{{2, 0}, {2, 1}, {1, 3}},
			},
			want: []int{1, 1, 0, 2},
		},
		{
			name: "example2",
			args: args{
				n:     3,
				edges: [][]int{{1, 2}, {2, 0}},
			},
			want: []int{2, 0, 1},
		},
		{
			name: "single_node",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "chain_reversals",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "star_outward",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}},
			},
			want: []int{0, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEdgeReversals(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minEdgeReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}
