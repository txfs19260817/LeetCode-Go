package leetcode

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0",
			args: args{2, 2, 2, 0, 0},
			want: 6,
		},
		{
			name: "m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1",
			args: args{1, 3, 3, 0, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
