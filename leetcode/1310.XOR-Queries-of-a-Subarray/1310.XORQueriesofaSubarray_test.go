package _310_XOR_Queries_of_a_Subarray

import (
	"reflect"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	type args struct {
		arr     []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]",
			args: args{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}},
			want: []int{2, 7, 14, 8},
		},
		{
			name: "arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]",
			args: args{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}},
			want: []int{8, 0, 4, 4},
		},
		{
			name: "arr = [16], queries = [[0,0],[0,0],[0,0]]",
			args: args{[]int{16}, [][]int{{0, 0}, {0, 0}, {0, 0}}},
			want: []int{16, 16, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorQueries(tt.args.arr, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
