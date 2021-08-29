package _588_Sum_of_All_Odd_Length_Subarrays

import "testing"

func Test_sumOddLengthSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [1,4,2,5,3]",
			args: args{[]int{1, 4, 2, 5, 3}},
			want: 58,
		},
		{
			name: "arr = [1,2]",
			args: args{[]int{1, 2}},
			want: 3,
		},
		{
			name: "arr = [10,11,12]",
			args: args{[]int{10, 11, 12}},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
