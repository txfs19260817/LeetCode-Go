package leetcode

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "12",
			args: args{12},
			want: 3,
		},
		{
			name: "13",
			args: args{13},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
			if got := numSquares1(tt.args.n); got != tt.want {
				t.Errorf("numSquares1() = %v, want %v", got, tt.want)
			}
			if got := numSquares2(tt.args.n); got != tt.want {
				t.Errorf("numSquares2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numSquares(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		numSquares(i)
	}
}

func Benchmark_numSquares1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		numSquares1(i)
	}
}

func Benchmark_numSquares2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		numSquares2(i)
	}
}
