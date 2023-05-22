package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "edges = [[1,2],[1,3],[2,3]]",
			args: args{[][]int{{1, 2}, {1, 3}, {2, 3}}},
			want: []int{2, 3},
		},
		{
			name: "edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]",
			args: args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
			want: []int{1, 4},
		},
		{
			name: "edges = [[3,4],[1,2],[2,4],[3,5],[2,5]]",
			args: args{[][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}}},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
