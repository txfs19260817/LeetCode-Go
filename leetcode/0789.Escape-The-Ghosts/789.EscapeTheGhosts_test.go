package _789_Escape_The_Ghosts

import "testing"

func Test_escapeGhosts(t *testing.T) {
	type args struct {
		ghosts [][]int
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ghosts = [[1,0],[0,3]], target = [0,1]",
			args: args{[][]int{{1, 0}, {0, 3}}, []int{0, 1}},
			want: true,
		},
		{
			name: "ghosts = [[5,0],[-10,-2],[0,-5],[-2,-2],[-7,1]], target = [7,7]",
			args: args{[][]int{{5, 0}, {-10, -2}, {0, -5}, {-2, -2}, {-7, 1}}, []int{7, 7}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeGhosts(tt.args.ghosts, tt.args.target); got != tt.want {
				t.Errorf("escapeGhosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
