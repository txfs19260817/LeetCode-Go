package leetcode

import "testing"

func Test_validTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			},
			want: false,
		},
		{
			name: "Single Node",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: true,
		},
		{
			name: "Two Nodes Disconnected",
			args: args{
				n:     2,
				edges: [][]int{},
			},
			want: false,
		},
		{
			name: "Two Nodes Connected",
			args: args{
				n:     2,
				edges: [][]int{{0, 1}},
			},
			want: true,
		},
		{
			name: "Cycle",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {1, 2}, {2, 0}},
			},
			want: false,
		},
		{
			name: "Disconnected Graph (Forest)",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {2, 3}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("validTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
