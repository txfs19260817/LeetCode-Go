package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNodesWithZeroOrOneParent(t *testing.T) {
	type args struct {
		inputs [][]int
	}
	tests := []struct {
		name   string
		args   args
		wantP0 []int
		wantP1 []int
	}{
		{
			name:   "1",
			args:   args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}},
			wantP0: []int{1, 2, 3},
			wantP1: []int{4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP0, gotP1 := findNodesWithZeroOrOneParent(tt.args.inputs)
			assert.ElementsMatch(t, gotP0, tt.wantP0)
			assert.ElementsMatch(t, gotP1, tt.wantP1)
		})
	}
}
