package leetcode

import "testing"

func Test_subarraySum(t *testing.T) {
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
			name: "nums = [1,1,1], k = 2",
			args: args{[]int{1, 1, 1}, 2},
			want: 2,
		},
		{
			name: "nums = [1,2,3], k = 3",
			args: args{[]int{1, 2, 3}, 3},
			want: 2,
		},
		{
			name: "nums = [0,0,0,0,0,0,0,0,0,0], k = 3",
			args: args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
