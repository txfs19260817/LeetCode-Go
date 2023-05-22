package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		args   args
		expect args
	}{
		{
			name: "",
			args: args{[][]int{
				{1, 0, 3, 4},
				{5, 0, 7, 0},
				{9, 1, 1, 1},
			}},
			expect: args{[][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{9, 0, 1, 0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			assert.EqualValues(t, tt.expect, tt.args)
		})
	}
}
