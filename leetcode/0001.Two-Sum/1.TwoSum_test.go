package _001_Two_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[2,7,11,15], 9",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}
}
