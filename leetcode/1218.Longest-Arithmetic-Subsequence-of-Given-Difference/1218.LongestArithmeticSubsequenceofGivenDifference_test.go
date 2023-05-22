package leetcode

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [1,2,3,4], difference = 1",
			args: args{[]int{1, 2, 3, 4}, 1},
			want: 4,
		},
		{
			name: "arr = [1,3,5,7], difference = 1",
			args: args{[]int{1, 3, 5, 7}, 1},
			want: 1,
		},
		{
			name: "arr = [1,5,7,8,5,3,4,2,1], difference = -2",
			args: args{[]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
