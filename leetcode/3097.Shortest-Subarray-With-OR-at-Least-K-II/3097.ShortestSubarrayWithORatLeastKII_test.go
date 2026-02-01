package leetcode

import "testing"

func Test_minimumSubarrayLength(t *testing.T) {
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
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3},
				k:    2,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 1, 8},
				k:    10,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2},
				k:    0,
			},
			want: 1,
		},
		{
			name: "Example 4 (Not Found)",
			args: args{
				nums: []int{1, 2},
				k:    10,
			},
			want: -1,
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{1, 2, 32, 21},
				k:    55,
			},
			want: 3,
		},
		{
			name: "Example 6",
			args: args{
				nums: []int{2, 1, 9, 12},
				k:    21,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
