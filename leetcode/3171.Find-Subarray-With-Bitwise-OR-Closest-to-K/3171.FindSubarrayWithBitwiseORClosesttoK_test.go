package leetcode

import "testing"

func Test_minimumDifference(t *testing.T) {
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
				nums: []int{1, 2, 4, 5},
				k:    3,
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 3, 1, 3},
				k:    2,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1},
				k:    10,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
