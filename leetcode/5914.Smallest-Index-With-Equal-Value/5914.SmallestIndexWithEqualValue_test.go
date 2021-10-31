package _914_Smallest_Index_With_Equal_Value

import "testing"

func Test_smallestEqual(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [0,1,2]",
			args: args{[]int{0, 1, 2}},
			want: 0,
		},
		{
			name: "nums = [4,3,2,1]",
			args: args{[]int{4, 3, 2, 1}},
			want: 2,
		},
		{
			name: "nums = [1,2,3,4,5,6,7,8,9,0]",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEqual(tt.args.nums); got != tt.want {
				t.Errorf("smallestEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
