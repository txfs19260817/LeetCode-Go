package _139_Largest_1_Bordered_Square

import "testing"

func Test_largest1BorderedSquare(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "grid = [[1,1,1],[1,0,1],[1,1,1]]",
			args: args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: 9,
		},
		{
			name: "grid = [[1,1,0,0]]",
			args: args{[][]int{{1, 1, 0, 0}}},
			want: 1,
		},
		{
			name: "grid = [[1,1,1,1],[1,0,0,1],[1,1,1,1]]",
			args: args{[][]int{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}}},
			want: 1,
		},
		{
			name: "grid = [[0,0,0,1]]",
			args: args{[][]int{{0, 0, 0, 1}}},
			want: 1,
		},
		{
			name: "grid = [[0],[0],[0],[1]]",
			args: args{[][]int{{0}, {0}, {0}, {1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largest1BorderedSquare(tt.args.grid); got != tt.want {
				t.Errorf("largest1BorderedSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
