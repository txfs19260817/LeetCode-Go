package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,2,0]",
			args: args{[]int{1, 2, 0}},
			want: 3,
		},
		{
			name: "nums = [3,4,-1,1]",
			args: args{[]int{3, 4, -1, 1}},
			want: 2,
		},
		{
			name: "nums = [7,8,9,11,12]",
			args: args{[]int{7, 8, 9, 11, 12}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
