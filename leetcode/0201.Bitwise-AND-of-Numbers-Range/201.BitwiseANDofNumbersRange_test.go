package leetcode

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "left = 5, right = 7",
			args: args{5, 7},
			want: 4,
		},
		{
			name: "left = 1, right = 2147483647",
			args: args{1, 2147483647},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
			if got := rangeBitwiseAnd2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_rangeBitwiseAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeBitwiseAnd(1, 2147483647)
	}
}

func Benchmark_rangeBitwiseAnd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeBitwiseAnd2(1, 2147483647)
	}
}
