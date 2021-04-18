package _087_Scramble_String

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s1 = "great", s2 = "rgeat"`,
			args: args{"great", "rgeat"},
			want: true,
		},
		{
			name: `s1 = "abcde", s2 = "caebd"`,
			args: args{"abcde", "caebd"},
			want: false,
		},
		{
			name: `s1 = "abcdbdacbdac", s2 = "bdacabcdbdac"`,
			args: args{"abcdbdacbdac", "bdacabcdbdac"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
