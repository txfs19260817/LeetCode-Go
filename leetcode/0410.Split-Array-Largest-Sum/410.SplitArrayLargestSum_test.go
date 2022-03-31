package _410_Split_Array_Largest_Sum

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [7,2,5,10,8], m = 2",
			args: args{[]int{7, 2, 5, 10, 8}, 2},
			want: 18,
		},
		{
			name: "nums = [1,4,4], m = 3",
			args: args{[]int{1, 4, 4}, 3},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
