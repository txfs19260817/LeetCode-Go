package leetcode

import "testing"

func Test_kthSmallestSubarraySum(t *testing.T) {
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
			name: "example 1",
			args: args{
				nums: []int{2, 1, 3},
				k:    4,
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{3, 3, 5, 5},
				k:    7,
			},
			want: 10,
		},
		{
			name: "single element",
			args: args{
				nums: []int{7},
				k:    1,
			},
			want: 7,
		},
		{
			name: "repeated sums",
			args: args{
				nums: []int{1, 1, 1},
				k:    5,
			},
			want: 2,
		},
		{
			name: "largest k returns total sum",
			args: args{
				nums: []int{2, 1, 3},
				k:    6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
