package _300_Longest_Increasing_Subsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [10,9,2,5,3,7,101,18]",
			args: args{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "nums = [0,1,0,3,2,3]",
			args: args{[]int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "nums = [7,7,7,7,7,7,7]",
			args: args{[]int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
		{
			name: "nums = [1,3,6,7,9,4,10,5,6]",
			args: args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
