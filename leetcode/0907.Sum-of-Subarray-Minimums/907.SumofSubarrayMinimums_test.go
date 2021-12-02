package _907_Sum_of_Subarray_Minimums

import "testing"

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [3,1,2,4]",
			args: args{[]int{3, 1, 2, 4}},
			want: 17,
		},
		{
			name: "arr = [11,81,94,43,3]",
			args: args{[]int{11, 81, 94, 43, 3}},
			want: 444,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.arr); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}
