package leetcode

import "testing"

func Test_minimumSemesters(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n:         3,
				relations: [][]int{{1, 3}, {2, 3}},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				n:         3,
				relations: [][]int{{1, 2}, {3, 1}, {2, 3}},
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				n:         4,
				relations: [][]int{{1, 2}, {3, 4}, {4, 3}},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSemesters(tt.args.n, tt.args.relations); got != tt.want {
				t.Errorf("minimumSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
