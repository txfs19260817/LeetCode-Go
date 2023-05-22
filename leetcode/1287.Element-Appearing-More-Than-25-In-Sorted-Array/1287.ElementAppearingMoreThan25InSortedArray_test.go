package leetcode

import "testing"

func Test_findSpecialInteger(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3,4,5,6,7,8,9,10,11,12,12,12,12]",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 12, 12}},
			want: 12,
		},
		{
			name: "[1,2,3,4,5,6,7,8,9]",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: -1,
		},
		{
			name: "[1,1]",
			args: args{[]int{1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSpecialInteger(tt.args.arr); got != tt.want {
				t.Errorf("findSpecialInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
