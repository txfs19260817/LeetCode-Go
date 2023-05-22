package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums  []int
		start int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,3], start = 6, goal = 4",
			args: args{[]int{1, 3}, 6, 4},
			want: 2,
		},
		{
			name: "nums = [2,4,12], start = 2, goal = 12",
			args: args{[]int{2, 4, 12}, 2, 12},
			want: 2,
		},
		{
			name: "nums = [3,5,7], start = 0, goal = -4",
			args: args{[]int{3, 5, 7}, 0, -4},
			want: 2,
		},
		{
			name: "nums = [2,8,16], start = 0, goal = 1",
			args: args{[]int{2, 8, 16}, 0, 1},
			want: -1,
		},
		{
			name: "nums = [1], start = 0, goal = 3",
			args: args{[]int{1}, 0, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums, tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
