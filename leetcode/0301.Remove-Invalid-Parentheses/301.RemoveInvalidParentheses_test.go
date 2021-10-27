package _301_Remove_Invalid_Parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "()())()",
			args: args{"()())()"},
			want: []string{"(())()", "()()()"},
		},
		{
			name: "(a)())()",
			args: args{"(a)())()"},
			want: []string{"(a())()", "(a)()()"},
		},
		{
			name: ")(",
			args: args{")("},
			want: []string{""},
		},
		{
			name: ")((((((((",
			args: args{")(((((((("},
			want: []string{""},
		},
		{
			name: "())))()v(k",
			args: args{"())))()v(k"},
			want: []string{"()()vk"},
		},
		{
			name: "))(p(((())",
			args: args{"))(p(((())"},
			want: []string{"p(())", "(p())"},
		},
		{
			name: ")()))())))",
			args: args{")()))())))"},
			want: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, removeInvalidParentheses(tt.args.s))
		})
	}
}
