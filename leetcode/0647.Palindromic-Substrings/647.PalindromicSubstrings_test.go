package _647_Palindromic_Substrings

import "testing"

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abc",
			args: args{"abc"},
			want: 3,
		},
		{
			name: "aaa",
			args: args{"aaa"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
			if got := countSubstrings2(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSubstrings("aaaaaaaaaaaaaaaaaaaaaaaaaaaaab")
	}
}

func Benchmark_countSubstrings2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSubstrings2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaab")
	}
}
