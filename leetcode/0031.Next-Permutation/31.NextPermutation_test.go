package _031_Next_Permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			name: "nums = [3,2,1]",
			args: args{[]int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "nums = [1,5,8,4,7,6,5,3,1]",
			args: args{[]int{1, 5, 8, 4, 7, 6, 5, 3, 1}},
			want: []int{1, 5, 8, 5, 1, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
