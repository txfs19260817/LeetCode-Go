package leetcode

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 1",
			args: args{1},
			want: 0,
		},
		{
			name: "n = 2",
			args: args{2},
			want: 1,
		},
		{
			name: "n = 3",
			args: args{3},
			want: 2,
		},
		{
			name: "n = 9999",
			args: args{9999},
			want: 24995000,
		},
		{
			name: "n = 10000",
			args: args{10000},
			want: 25000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
