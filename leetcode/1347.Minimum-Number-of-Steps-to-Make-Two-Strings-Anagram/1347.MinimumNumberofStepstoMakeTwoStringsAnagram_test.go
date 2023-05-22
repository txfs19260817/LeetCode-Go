package leetcode

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "bab", t = "aba"`,
			args: args{"bab", "aba"},
			want: 1,
		},
		{
			name: `s = "xxyyzz", t = "xxyyzz"`,
			args: args{"xxyyzz", "xxyyzz"},
			want: 0,
		},
		{
			name: `s = "friend", t = "family"`,
			args: args{"friend", "family"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
