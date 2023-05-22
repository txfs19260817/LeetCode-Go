package leetcode

import "testing"

func Test_makeConnected(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 4, connections = [[0,1],[0,2],[1,2]]",
			args: args{4, [][]int{{0, 1}, {0, 2}, {1, 2}}},
			want: 1,
		},
		{
			name: "n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]",
			args: args{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}},
			want: 2,
		},
		{
			name: "n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]",
			args: args{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}},
			want: -1,
		},
		{
			name: "n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]",
			args: args{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
