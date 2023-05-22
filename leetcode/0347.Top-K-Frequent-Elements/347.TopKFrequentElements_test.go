package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,1,1,2,2,3], k = 2",
			args: args{[]int{1, 1, 1, 2, 2, 3}, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, topKFrequent(tt.args.nums, tt.args.k))
		})
	}
}
