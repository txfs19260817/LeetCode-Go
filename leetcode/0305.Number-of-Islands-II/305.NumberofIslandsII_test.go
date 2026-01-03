package leetcode

import (
	"reflect"
	"testing"
)

func Test_numIslands2(t *testing.T) {
	type args struct {
		m         int
		n         int
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}},
			},
			want: []int{1, 1, 2, 3},
		},
		{
			name: "Example with duplicates",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {0, 0}},
			},
			want: []int{1, 1, 2, 3, 3},
		},
		{
			name: "Single island",
			args: args{
				m:         1,
				n:         1,
				positions: [][]int{{0, 0}},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(tt.args.m, tt.args.n, tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}
