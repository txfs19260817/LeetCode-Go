package leetcode

import "testing"

func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{6, 0, 8, 2, 1, 5},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1},
			},
			want: 7,
		},
		{
			name: "Long Decreasing Sequence With Duplicate Tail",
			args: args{
				nums: func() []int {
					nums := make([]int, 0, 50000-10+2)
					for v := 50000; v >= 10; v-- {
						nums = append(nums, v)
					}
					nums = append(nums, 10)
					return nums
				}(),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.nums); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
