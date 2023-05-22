package leetcode

import (
	"reflect"
	"testing"
)

func Test_countRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
		points     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "rectangles = [[1,1],[2,2],[3,3]], points = [[1,3],[1,1]]",
			args: args{[][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}}},
			want: []int{1, 3},
		},
		{
			name: "rectangles = [[7,1],[2,6],...], points = [[10,3],[8,10],...",
			args: args{[][]int{{7, 1}, {2, 6}, {1, 4}, {5, 2}, {10, 3}, {2, 4}, {5, 9}}, [][]int{{10, 3}, {8, 10}, {2, 3}, {5, 4}, {8, 5}, {7, 10}, {6, 6}, {3, 6}}},
			want: []int{1, 0, 4, 1, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRectangles(tt.args.rectangles, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
