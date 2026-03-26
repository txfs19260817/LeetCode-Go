package leetcode

import "testing"

func Test_findShortestCycle(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     7,
				edges: [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {0, 2}},
			},
			want: -1,
		},
		{
			name: "Disconnected Graph With Triangle Component",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "Square Cycle",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			},
			want: 4,
		},
		{
			name: "Multiple Cycles Returns Shortest",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}, {4, 5}, {5, 1}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestCycle(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("findShortestCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
