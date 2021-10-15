package _542_01_Matrix

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "mat = [[0,0,0],[0,1,0],[0,0,0]]",
			args: args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name: "mat = [[0,0,0],[0,1,0],[1,1,1]]",
			args: args{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
			if got := updateMatrix2(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
