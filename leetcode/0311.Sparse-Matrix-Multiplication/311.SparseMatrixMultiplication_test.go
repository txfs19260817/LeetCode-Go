package leetcode

import (
	"reflect"
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		mat1 [][]int
		mat2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example",
			args: args{
				mat1: [][]int{
					{1, 0, 0},
					{-1, 0, 3},
				},
				mat2: [][]int{
					{7, 0, 0},
					{0, 0, 0},
					{0, 0, 1},
				},
			},
			want: [][]int{
				{7, 0, 0},
				{-7, 0, 3},
			},
		},
		{
			name: "all zero left matrix",
			args: args{
				mat1: [][]int{
					{0, 0},
					{0, 0},
				},
				mat2: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "rectangular result",
			args: args{
				mat1: [][]int{
					{1, 2},
					{3, 4},
				},
				mat2: [][]int{
					{5},
					{6},
				},
			},
			want: [][]int{
				{17},
				{39},
			},
		},
		{
			name: "single element",
			args: args{
				mat1: [][]int{{-2}},
				mat2: [][]int{{3}},
			},
			want: [][]int{{-6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.mat1, tt.args.mat2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
