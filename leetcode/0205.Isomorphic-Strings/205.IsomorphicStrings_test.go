package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s = "egg", t = "add"`,
			args: args{"egg", "add"},
			want: true,
		},
		{
			name: `s = "paper", t = "title"`,
			args: args{"paper", "title"},
			want: true,
		},
		{
			name: `s = "foo", t = "bar"`,
			args: args{"foo", "bar"},
			want: false,
		},
		{
			name: `s = "bbbaaaba", t = "aaabbbba"`,
			args: args{"bbbaaaba", "aaabbbba"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
