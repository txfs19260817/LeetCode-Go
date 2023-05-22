package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `S = "a1b2"`,
			args: args{"a1b2"},
			want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			name: `S = "3z4"`,
			args: args{"3z4"},
			want: []string{"3z4", "3Z4"},
		},
		{
			name: `S = "12345"`,
			args: args{"12345"},
			want: []string{"12345"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, letterCasePermutation(tt.args.S))
			assert.ElementsMatch(t, tt.want, letterCasePermutation1(tt.args.S))
		})
	}
}

func Benchmark_letterCasePermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = letterCasePermutation("xxulYjUwme")
	}
}

func Benchmark_letterCasePermutation1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = letterCasePermutation1("xxulYjUwme")
	}
}
