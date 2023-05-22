package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `pattern = "abba", s = "dog cat cat fish"`,
			args: args{"abba", "dog cat cat fish"},
			want: false,
		},
		{
			name: `pattern = "aaa", s = "aa aa aa aa"`,
			args: args{"aaa", "aa aa aa aa"},
			want: false,
		},
		{
			name: `pattern = "abba", s = "dog dog dog dog"`,
			args: args{"abba", "dog dog dog dog"},
			want: false,
		},
		{
			name: `pattern = "abba", s = "dog cat cat dog"`,
			args: args{"abba", "dog cat cat dog"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
