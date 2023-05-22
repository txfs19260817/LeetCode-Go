package leetcode

import "testing"

func Test_getMaxLen(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,-2,-3,4]",
			args: args{[]int{1, -2, -3, 4}},
			want: 4,
		},
		{
			name: "[0,1,-2,-3,-4]",
			args: args{[]int{0, 1, -2, -3, -4}},
			want: 3,
		},
		{
			name: "[-1,-2,-3,0,1]",
			args: args{[]int{-1, -2, -3, 0, 1}},
			want: 2,
		},
		{
			name: "[-1,2]",
			args: args{[]int{-1, 2}},
			want: 1,
		},
		{
			name: "[1,2,3,5,-6,4,0,10]",
			args: args{[]int{1, 2, 3, 5, -6, 4, 0, 10}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.args.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
