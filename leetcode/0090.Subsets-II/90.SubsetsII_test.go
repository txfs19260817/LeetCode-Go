package _090_Subsets_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,2,2]",
			args: args{[]int{1, 2, 2}},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, subsetsWithDup(tt.args.nums), tt.want)
		})
	}
}
