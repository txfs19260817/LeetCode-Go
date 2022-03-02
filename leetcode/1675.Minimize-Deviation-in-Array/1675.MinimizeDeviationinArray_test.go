package _675_Minimize_Deviation_in_Array

import "testing"

func Test_minimumDeviation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3,4]",
			args: args{[]int{1, 2, 3, 4}},
			want: 1,
		},
		{
			name: "[4,1,5,20,3]",
			args: args{[]int{4, 1, 5, 20, 3}},
			want: 3,
		},
		{
			name: "[2,10,8]",
			args: args{[]int{2, 10, 8}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeviation(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
