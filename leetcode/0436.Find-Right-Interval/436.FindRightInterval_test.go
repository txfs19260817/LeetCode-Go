package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "intervals = [[1,2]]",
			args: args{[][]int{{1, 2}}},
			want: []int{-1},
		},
		{
			name: "intervals = [[3,4],[2,3],[1,2]]",
			args: args{[][]int{{3, 4}, {2, 3}, {1, 2}}},
			want: []int{-1, 0, 1},
		},
		{
			name: "intervals = [[1,4],[2,3],[3,4]]",
			args: args{[][]int{{1, 4}, {2, 3}, {3, 4}}},
			want: []int{-1, 2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightInterval(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
