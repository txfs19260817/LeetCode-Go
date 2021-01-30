package _424_Longest_Repeating_Character_Replacement

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "s = `ABAB`, k = 2",
			args: args{"ABAB", 2},
			want: 4,
		},
		{
			name: "s = `AABABBA`, k = 1",
			args: args{"AABABBA", 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement1(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_characterReplacement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		characterReplacement("AABABBA", 1)
	}
}

func Benchmark_characterReplacement1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		characterReplacement1("AABABBA", 1)
	}
}
