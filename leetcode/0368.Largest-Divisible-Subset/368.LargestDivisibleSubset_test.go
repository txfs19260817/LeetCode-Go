package _368_Largest_Divisible_Subset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: [][]int{{1, 2}, {1, 3}},
		},
		{
			name: "nums = [1,2,4,8]",
			args: args{[]int{1, 2, 4, 8}},
			want: [][]int{{1, 2, 4, 8}},
		},
		{
			name: "nums = [4,8,10,240]",
			args: args{[]int{4, 8, 10, 240}},
			want: [][]int{{4, 8, 240}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.want, largestDivisibleSubset(tt.args.nums))
		})
	}
}
