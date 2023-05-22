package leetcode

import (
	"reflect"
	"testing"
)

func Test_getModifiedArray(t *testing.T) {
	type args struct {
		length  int
		updates [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]",
			args: args{5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}},
			want: []int{-2, 0, 3, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getModifiedArray(tt.args.length, tt.args.updates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModifiedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
