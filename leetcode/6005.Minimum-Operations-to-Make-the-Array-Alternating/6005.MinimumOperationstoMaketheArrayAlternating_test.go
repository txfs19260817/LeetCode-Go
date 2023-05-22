package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,1,3,2,4,3]",
			args: args{[]int{3, 1, 3, 2, 4, 3}},
			want: 3,
		},
		{
			name: "[1,2,2,2,2]",
			args: args{[]int{1, 2, 2, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
