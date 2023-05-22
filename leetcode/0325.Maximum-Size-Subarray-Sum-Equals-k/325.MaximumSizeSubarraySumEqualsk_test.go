package leetcode

import "testing"

func Test_maxSubArrayLen(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,-1,5,-2,3], k = 3",
			args: args{[]int{1, -1, 5, -2, 3}, 3},
			want: 4,
		},
		{
			name: "nums = [-2,-1,2,1], k = 1",
			args: args{[]int{-2, -1, 2, 1}, 1},
			want: 2,
		},
		{
			name: "nums = [1,1,0], k = 1",
			args: args{[]int{1, 1, 0}, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayLen(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
