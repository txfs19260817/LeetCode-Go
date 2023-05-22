package leetcode

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "125",
			args: args{125},
			want: true,
		},
		{
			name: "1",
			args: args{1},
			want: true,
		},
		{
			name: "16",
			args: args{16},
			want: true,
		},
		{
			name: "46",
			args: args{46},
			want: true,
		},
		{
			name: "10",
			args: args{10},
			want: false,
		},
		{
			name: "24",
			args: args{24},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
			if got := reorderedPowerOf22(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf22() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reorderedPowerOf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reorderedPowerOf2(b.N)
	}
}

func Benchmark_reorderedPowerOf22(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reorderedPowerOf22(b.N)
	}
}
