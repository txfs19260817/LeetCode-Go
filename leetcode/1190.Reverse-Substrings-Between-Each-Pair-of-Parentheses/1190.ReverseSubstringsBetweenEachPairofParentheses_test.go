package leetcode

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "(abcd)",
			args: args{"(abcd)"},
			want: "dcba",
		},
		{
			name: "(x(ab)y(cd)z)",
			args: args{"(x(ab)y(cd)z)"},
			want: "zcdyabx",
		},
		{
			name: "(uoy(love)mi(hate)ko)",
			args: args{"(uoy(love)mi(hate)ko)"},
			want: "okhateimloveyou",
		},
		{
			name: "a(bcdefghijkl(mno)p)q",
			args: args{"a(bcdefghijkl(mno)p)q"},
			want: "apmnolkjihgfedcbq",
		},
		{
			name: "a",
			args: args{"a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
