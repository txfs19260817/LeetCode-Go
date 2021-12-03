package k

import "testing"

func Test_hasCommonAncestor(t *testing.T) {
	type args struct {
		edges [][]int
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}, 4, 5},
			want: true,
		},
		{
			name: "2",
			args: args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}, 1, 2},
			want: false,
		},
		{
			name: "3",
			args: args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}, 3, 2},
			want: false,
		},
		{
			name: "4",
			args: args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}, 6, 7},
			want: true,
		},
		{
			name: "5",
			args: args{[][]int{{1, 4}, {1, 5}, {2, 5}, {3, 6}, {6, 7}}, 1, 7},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCommonAncestor(tt.args.edges, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hasCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
