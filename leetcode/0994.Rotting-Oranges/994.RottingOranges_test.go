package _994_Rotting_Oranges

import "testing"

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "grid = [[2,1,1],[1,1,0],[0,1,1]]",
			args: args{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}},
			want: 4,
		},
		{
			name: "grid = [[2,1,1],[0,1,1],[1,0,1]]",
			args: args{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}},
			want: -1,
		},
		{
			name: "grid = [[0,2]]",
			args: args{[][]int{{0, 2}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
