package _260_Single_Number_III

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,1,3,2,5]",
			args: args{[]int{1, 2, 1, 3, 2, 5}},
			want: []int{3, 5},
		},
		{
			name: "nums = [1,0]",
			args: args{[]int{1, 0}},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, singleNumber(tt.args.nums))
		})
	}
}
