package leetcode

import "testing"

func Test_champagneTower(t *testing.T) {
	type args struct {
		poured      int
		query_row   int
		query_glass int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "poured = 1, query_row = 1, query_glass = 1",
			args: args{1, 1, 1},
			want: 0,
		},
		{
			name: "poured = 100000009, query_row = 33, query_glass = 17",
			args: args{100000009, 33, 17},
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := champagneTower(tt.args.poured, tt.args.query_row, tt.args.query_glass); got != tt.want {
				t.Errorf("champagneTower() = %v, want %v", got, tt.want)
			}
		})
	}
}
