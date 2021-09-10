package _413_Arithmetic_Slices

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,2,3,4]",
			args: args{[]int{1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "nums = [2,4,6,8,9,10,11]",
			args: args{[]int{2, 4, 6, 8, 9, 10, 11}},
			want: 6,
		},
		{
			name: "nums = [1]",
			args: args{[]int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
