package leetcode

import (
	"reflect"
	"testing"
)

func Test_matrixRankTransform(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "matrix = [[1,2],[3,4]]",
			args: args{[][]int{{1, 2}, {3, 4}}},
			want: [][]int{{1, 2}, {2, 3}},
		},
		{
			name: "matrix = [[7,7],[7,7]]",
			args: args{[][]int{{7, 7}, {7, 7}}},
			want: [][]int{{1, 1}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixRankTransform(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
