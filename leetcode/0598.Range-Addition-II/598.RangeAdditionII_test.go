package leetcode

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m = 3, n = 3, ops = [[2,2],[3,3]]",
			args: args{3, 3, [][]int{{2, 2}, {3, 3}}},
			want: 4,
		},
		{
			name: "m = 3, n = 3, ops = []",
			args: args{3, 3, [][]int{}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
