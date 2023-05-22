package leetcode

import (
	"reflect"
	"testing"
)

func Test_arrayChange(t *testing.T) {
	type args struct {
		nums       []int
		operations [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,4,6], operations = [[1,3],[4,7],[6,1]]",
			args: args{[]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}},
			want: []int{3, 2, 7, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayChange(tt.args.nums, tt.args.operations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
