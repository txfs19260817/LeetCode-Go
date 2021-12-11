package sn

import "testing"

func Test_largestSubGrid(t *testing.T) {
	type args struct {
		grid   [][]int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{[][]int{{2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 3},
			want: 0,
		},
		{
			name: "1",
			args: args{[][]int{{2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 4},
			want: 1,
		},
		{
			name: "2",
			args: args{[][]int{{2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 14},
			want: 2,
		},
		{
			name: "3",
			args: args{[][]int{{2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 27},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubGrid(tt.args.grid, tt.args.maxSum); got != tt.want {
				t.Errorf("largestSubGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
