package leetcode

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2736",
			args: args{2736},
			want: 7236,
		},
		{
			name: "9973",
			args: args{9973},
			want: 9973,
		},
		{
			name: "1992",
			args: args{1992},
			want: 9912,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
			if got := maximumSwap2(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maximumSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumSwap(i)
	}
}

func Benchmark_maximumSwap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumSwap2(i)
	}
}
