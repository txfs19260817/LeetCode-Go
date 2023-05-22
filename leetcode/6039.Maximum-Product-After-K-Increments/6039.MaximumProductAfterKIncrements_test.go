package leetcode

import "testing"

func Test_maximumProduct(t *testing.T) {
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
			name: "nums = [0,4], k = 5",
			args: args{[]int{0, 4}, 5},
			want: 20,
		},
		{
			name: "nums = [6,3,3,2], k = 2",
			args: args{[]int{6, 3, 3, 2}, 2},
			want: 216,
		},
		{
			name: "nums = [9,7,8], k = 9",
			args: args{[]int{9, 7, 8}, 9},
			want: 1331,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
