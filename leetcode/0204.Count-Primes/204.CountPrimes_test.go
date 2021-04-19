package _204_Count_Primes

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{10},
			want: 4,
		},
		{
			name: "0",
			args: args{0},
			want: 0,
		},
		{
			name: "1",
			args: args{1},
			want: 0,
		},
		{
			name: "2",
			args: args{2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
