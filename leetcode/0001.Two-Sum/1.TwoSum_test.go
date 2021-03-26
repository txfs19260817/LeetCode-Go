package _001_Two_Sum

import (
	"reflect"
	"testing"

	"github.com/txfs19260817/leetcode-go/leetcode/utils"
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
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) && !utils.SameSlice(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
