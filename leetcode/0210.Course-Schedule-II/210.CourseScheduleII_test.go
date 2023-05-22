package leetcode

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]",
			args: args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "numCourses = 7, prerequisites = [[1,0],[0,3],[0,2],[3,2],[2,5],[4,5],[5,6],[2,4]]",
			args: args{7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}},
			want: []int{6, 5, 4, 2, 3, 0, 1},
		},
		{
			name: "numCourses = 2, prerequisites = [[1,0],[0,1]]",
			args: args{2, [][]int{{1, 0}, {0, 1}}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
