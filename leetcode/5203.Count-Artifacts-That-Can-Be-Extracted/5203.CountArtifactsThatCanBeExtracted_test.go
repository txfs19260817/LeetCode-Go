package leetcode

import "testing"

func Test_digArtifacts(t *testing.T) {
	type args struct {
		n         int
		artifacts [][]int
		dig       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2, artifacts = [[0,0,0,0],[0,1,1,1]], dig = [[0,0],[0,1]]",
			args: args{2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digArtifacts(tt.args.n, tt.args.artifacts, tt.args.dig); got != tt.want {
				t.Errorf("digArtifacts() = %v, want %v", got, tt.want)
			}
		})
	}
}
