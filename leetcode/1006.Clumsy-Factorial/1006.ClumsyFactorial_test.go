package leetcode

import "testing"

func Test_clumsy(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{4},
			want: 7,
		},
		{
			name: "10",
			args: args{10},
			want: 12,
		},
		{
			name: "3",
			args: args{3},
			want: 6,
		},
		{
			name: "2",
			args: args{2},
			want: 2,
		},
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "10000",
			args: args{10000},
			want: 10001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.args.N); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
