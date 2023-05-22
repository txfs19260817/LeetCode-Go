package leetcode

import "testing"

func Test_tilingRectangle(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2, m = 3",
			args: args{2, 3},
			want: 3,
		},
		{
			name: "n = 1, m = 1",
			args: args{1, 1},
			want: 1,
		},
		{
			name: "n = 11, m = 13",
			args: args{11, 13},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tilingRectangle(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("tilingRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
