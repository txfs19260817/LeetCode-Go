package leetcode

import "testing"

func Test_numDistinctIslands(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{
					{1, 1, 0, 1, 1},
					{1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1},
					{1, 1, 0, 1, 1},
				},
			},
			want: 3,
		},
		{
			name: "Single Cell",
			args: args{
				grid: [][]int{
					{1},
				},
			},
			want: 1,
		},
		{
			name: "Empty Grid",
			args: args{
				grid: [][]int{
					{0, 0},
					{0, 0},
				},
			},
			want: 0,
		},
		{
			name: "Diagonal Distinct",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			// {1,0} island: path 0->end
			// {0,1} island: path 0->end
			// They are same shape?
			// Wait, the shape encoding needs to be relative to start point.
			// Path encoding is relative to DFS traversal.
			// With canonical path encoding, they should be the same.
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinctIslands(tt.args.grid); got != tt.want {
				t.Errorf("numDistinctIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
