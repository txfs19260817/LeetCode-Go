package leetcode

import (
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1",
			args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1},
			want: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			name: "mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2",
			args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2},
			want: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixBlockSum(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
