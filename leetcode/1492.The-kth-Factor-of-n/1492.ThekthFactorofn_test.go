package leetcode

import "testing"

func Test_kthFactor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 12, k = 3",
			args: args{12, 3},
			want: 3,
		},
		{
			name: "n = 7, k = 2",
			args: args{7, 2},
			want: 7,
		},
		{
			name: "n = 4, k = 4",
			args: args{4, 4},
			want: -1,
		},
		{
			name: "n = 1, k = 1",
			args: args{1, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthFactor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
