package leetcode

import "testing"

func Test_xorGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [1, 1, 2]",
			args: args{[]int{1, 1, 2}},
			want: false,
		},
		{
			name: "nums = [1, 1]",
			args: args{[]int{1, 1}},
			want: true,
		},
		{
			name: "nums = [1, 2, 3]",
			args: args{[]int{1, 2, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorGame(tt.args.nums); got != tt.want {
				t.Errorf("xorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
