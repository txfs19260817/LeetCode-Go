package k

import "testing"

func Test_nonogram(t *testing.T) {
	type args struct {
		mat  [][]int
		rows [][]int
		cols [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {0, 1, 0, 0}, {1, 1, 0, 1}, {0, 0, 1, 1}}, [][]int{{}, {1}, {1, 2}, {1}, {2}}, [][]int{{2, 1}, {1}, {2}, {1}}},
			want: true,
		},
		{
			name: "2",
			args: args{[][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {0, 1, 0, 0}, {1, 1, 0, 1}, {0, 0, 1, 1}}, [][]int{{}, {}, {1}, {1}, {1, 1}}, [][]int{{2}, {1}, {2}, {1}}},
			want: false,
		},
		{
			name: "3",
			args: args{[][]int{{1, 1}, {0, 0}, {0, 0}, {1, 0}}, [][]int{{}, {2}, {2}, {1}}, [][]int{{1, 1}, {3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonogram(tt.args.mat, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("nonogram() = %v, want %v", got, tt.want)
			}
		})
	}
}
