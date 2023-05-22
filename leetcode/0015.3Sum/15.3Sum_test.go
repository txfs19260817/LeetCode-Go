package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "0, 0, 0",
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "-1, 0, 1, 2, -1, -4",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6",
			args: args{nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
		{
			name: "5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5",
			args: args{nums: []int{5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5}},
			want: [][]int{{-10, 2, 8}, {-10, 3, 7}, {-10, 4, 6}, {-10, 5, 5}, {-8, 2, 6}, {-8, 3, 5}, {-8, 4, 4}, {-7, -1, 8},
				{-7, 2, 5}, {-7, 3, 4}, {-6, -2, 8}, {-6, -1, 7}, {-6, 2, 4}, {-5, -3, 8}, {-5, -2, 7}, {-5, -1, 6}, {-5, 2, 3},
				{-4, -3, 7}, {-4, -2, 6}, {-4, -1, 5}, {-4, 2, 2}, {-3, -3, 6}, {-3, -2, 5}, {-3, -1, 4}, {-2, -1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_threeSum(b *testing.B) {
	input := []int{5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5}
	for i := 0; i < b.N; i++ {
		_ = threeSum(input)
	}
}
