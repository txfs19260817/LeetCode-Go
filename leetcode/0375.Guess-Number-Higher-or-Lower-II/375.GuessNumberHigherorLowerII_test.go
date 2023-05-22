package leetcode

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{4},
			want: 4,
		},
		{
			name: "5",
			args: args{5},
			want: 6,
		},
		{
			name: "6",
			args: args{6},
			want: 8,
		},
		{
			name: "200",
			args: args{200},
			want: 952,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
			if got := getMoneyAmount2(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getMoneyAmount(b *testing.B) {
	for i := 1; i <= 200; i++ {
		getMoneyAmount(i)
	}
}

func Benchmark_getMoneyAmount2(b *testing.B) {
	for i := 1; i <= 200; i++ {
		getMoneyAmount2(i)
	}
}
