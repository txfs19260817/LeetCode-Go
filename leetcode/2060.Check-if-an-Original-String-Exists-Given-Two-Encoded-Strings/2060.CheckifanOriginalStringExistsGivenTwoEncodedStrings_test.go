package leetcode

import "testing"

func Test_possiblyEquals(t *testing.T) {
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
			name: `s1 = "internationalization", s2 = "i18n"`,
			args: args{"internationalization", "i18n"},
			want: true,
		},
		{
			name: `s1 = "l123e", s2 = "44"`,
			args: args{"l123e", "44"},
			want: true,
		},
		{
			name: `s1 = "112s", s2 = "g841"`,
			args: args{"112s", "g841"},
			want: true,
		},
		{
			name: `s1 = "a5b", s2 = "c5b"`,
			args: args{"a5b", "c5b"},
			want: false,
		},
		{
			name: `s1 = "ab", s2 = "a2"`,
			args: args{"ab", "a2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possiblyEquals(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("possiblyEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}
