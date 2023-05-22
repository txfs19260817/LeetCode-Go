package leetcode

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: 3,
		},
		{
			name: "[1,1,1]",
			args: args{[]int{1, 1, 1}},
			want: 0,
		},
		{
			name: "[1,1,1000000000]",
			args: args{[]int{1, 1, 1000000000}},
			want: 999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
