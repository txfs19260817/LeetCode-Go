package _153_Find_Minimum_in_Rotated_Sorted_Array

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,4,5,1,2]",
			args: args{[]int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "nums = [4,5,6,7,0,1,2]",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
		{
			name: "nums = [11,13,15,17]",
			args: args{[]int{11, 13, 15, 17}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
