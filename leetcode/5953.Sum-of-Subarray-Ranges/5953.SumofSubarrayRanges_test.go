package leetcode

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "[1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: 4,
		},
		{
			name: "[1,3,3]",
			args: args{[]int{1, 3, 3}},
			want: 4,
		},
		{
			name: "[4,-2,-3,4,1]",
			args: args{[]int{4, -2, -3, 4, 1}},
			want: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
