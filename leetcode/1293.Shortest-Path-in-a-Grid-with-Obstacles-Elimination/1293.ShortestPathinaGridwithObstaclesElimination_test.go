package leetcode

import "testing"

func Test_shortestPath(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0},
			},
			k:    1,
			want: 6,
		},
		{
			name: "example 2",
			grid: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
			k:    1,
			want: -1,
		},
		{
			name: "no obstacles",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			k:    0,
			want: 4,
		},
		{
			name: "one elimination needed",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			k:    1,
			want: 2,
		},
		{
			name: "not enough elimination",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			k:    0,
			want: -1,
		},
		{
			name: "start equals end",
			grid: [][]int{
				{0},
			},
			k:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := shortestPath(tt.grid, tt.k); got != tt.want {
				t.Fatalf("shortestPath() = %d, want %d", got, tt.want)
			}
		})
	}
}
