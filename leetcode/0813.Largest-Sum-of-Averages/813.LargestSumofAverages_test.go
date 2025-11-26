package leetcode

import (
	"math"
	"testing"
)

func Test_largestSumOfAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example1",
			args: args{nums: []int{9, 1, 2, 3, 9}, k: 3},
			want: 20.0,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 4},
			want: 20.5,
		},
		{
			name: "single_element",
			args: args{nums: []int{5}, k: 1},
			want: 5.0,
		},
		{
			name: "k_equals_1",
			args: args{nums: []int{1, 2, 3, 4}, k: 1},
			want: 2.5,
		},
		{
			name: "k_equals_len",
			args: args{nums: []int{2, 2, 2}, k: 3},
			want: 6.0,
		},
		{
			name: "all_zeros",
			args: args{nums: []int{0, 0, 0}, k: 2},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumOfAverages(tt.args.nums, tt.args.k); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("largestSumOfAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
