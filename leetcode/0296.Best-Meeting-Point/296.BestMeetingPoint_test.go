package leetcode

import "testing"

func Test_minTotalDistance(t *testing.T) {
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
					{1, 0, 0, 0, 1},
					{0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0},
				},
			},
			want: 6,
		},
		{
			name: "200x200 all 1s",
			args: args{
				grid: makeAllOnesGrid(200, 200),
			},
			want: 4000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTotalDistance(tt.args.grid); got != tt.want {
				t.Errorf("minTotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeAllOnesGrid(m, n int) [][]int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			grid[i][j] = 1
		}
	}
	return grid
}
