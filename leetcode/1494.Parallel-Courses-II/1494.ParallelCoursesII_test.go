package leetcode

import "testing"

func Test_minNumberOfSemesters(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 4, dependencies = [[2,1],[3,1],[1,4]], k = 2",
			args: args{4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2},
			want: 3,
		},
		{
			name: "n = 5, dependencies = [[2,1],[3,1],[4,1],[1,5]], k = 2",
			args: args{5, [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}, 2},
			want: 4,
		},
		{
			name: "n = 11, dependencies = [], k = 2",
			args: args{11, [][]int{}, 2},
			want: 6,
		},
		{
			name: "n = 5, dependencies = [[1,5],[1,3],[1,2],[4,2],[4,5],[2,5],[1,4],[4,3],[3,5],[3,2]], k = 3",
			args: args{5, [][]int{{1, 5}, {1, 3}, {1, 2}, {4, 2}, {4, 5}, {2, 5}, {1, 4}, {4, 3}, {3, 5}, {3, 2}}, 3},
			want: 5,
		},
		{
			name: "n = 13, dependencies = [[12,8],[2,4],[3,7],[6,8],[11,8],[9,4],[9,7],[12,4],[11,4],[6,4],[1,4],[10,7],[10,4],[1,7],[1,8],[2,7],[8,4],[10,8],[12,7],[5,4],[3,4],[11,7],[7,4],[13,4],[9,8],[13,8]], k = 9",
			args: args{13, [][]int{{12, 8}, {2, 4}, {3, 7}, {6, 8}, {11, 8}, {9, 4}, {9, 7}, {12, 4}, {11, 4}, {6, 4}, {1, 4}, {10, 7}, {10, 4}, {1, 7}, {1, 8}, {2, 7}, {8, 4}, {10, 8}, {12, 7}, {5, 4}, {3, 4}, {11, 7}, {7, 4}, {13, 4}, {9, 8}, {13, 8}}, 9},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfSemesters(tt.args.n, tt.args.relations, tt.args.k); got != tt.want {
				t.Errorf("minNumberOfSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
