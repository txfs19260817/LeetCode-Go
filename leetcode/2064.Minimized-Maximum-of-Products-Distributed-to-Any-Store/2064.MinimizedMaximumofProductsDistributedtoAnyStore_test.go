package leetcode

import "testing"

func Test_minimizedMaximum(t *testing.T) {
	type args struct {
		n          int
		quantities []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 6, quantities = [11,6]",
			args: args{6, []int{11, 6}},
			want: 3,
		},
		{
			name: "n = 7, quantities = [15,10,10]",
			args: args{7, []int{15, 10, 10}},
			want: 5,
		},
		{
			name: "n = 1, quantities = [1]",
			args: args{1, []int{1}},
			want: 1,
		},
		{
			name: "n = 1, quantities = [100000]",
			args: args{1, []int{100000}},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizedMaximum(tt.args.n, tt.args.quantities); got != tt.want {
				t.Errorf("minimizedMaximum() = %v, want %v", got, tt.want)
			}
			if got := minimizedMaximum2(tt.args.n, tt.args.quantities); got != tt.want {
				t.Errorf("minimizedMaximum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minimizedMaximum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimizedMaximum(1, []int{100000})
	}
}

func Benchmark_minimizedMaximum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimizedMaximum2(1, []int{100000})
	}
}
