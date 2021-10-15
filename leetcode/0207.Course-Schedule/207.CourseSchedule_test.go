package _207_Course_Schedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "numCourses = 2, prerequisites = [[1,0]]",
			args: args{2, [][]int{{1, 0}}},
			want: true,
		},
		{
			name: "numCourses = 2, prerequisites = [[1,0],[0,1]]",
			args: args{2, [][]int{{1, 0}, {0, 1}}},
			want: false,
		},
		{
			name: "numCourses = 5, prerequisites = [[1,4],[2,4],[3,1],[3,2]]",
			args: args{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
