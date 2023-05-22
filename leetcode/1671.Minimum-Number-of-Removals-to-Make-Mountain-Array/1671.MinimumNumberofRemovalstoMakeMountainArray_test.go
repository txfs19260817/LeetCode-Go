package leetcode

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,1]",
			args: args{[]int{1, 3, 1}},
			want: 0,
		},
		{
			name: "[2,1,1,5,6,2,3,1]",
			args: args{[]int{2, 1, 1, 5, 6, 2, 3, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.args.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
