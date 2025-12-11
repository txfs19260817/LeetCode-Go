package leetcode

import "testing"

func Test_maxNumEdgesToRemove(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     2,
				edges: [][]int{{1, 1, 2}, {2, 1, 2}, {3, 1, 2}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				n:     4,
				edges: [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}},
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				n:     4,
				edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}},
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				n:     4,
				edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumEdgesToRemove(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("maxNumEdgesToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
