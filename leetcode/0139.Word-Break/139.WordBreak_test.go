package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wordBreakBenchResult bool

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, wordBreak(tt.args.s, tt.args.wordDict), tt.want)
			assert.Equal(t, wordBreak2(tt.args.s, tt.args.wordDict), tt.want)
		})
	}
}

func Benchmark_wordBreak(b *testing.B) {
	benchmarks := []struct {
		name     string
		s        string
		wordDict []string
	}{
		{
			name:     "example_1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
		},
		{
			name:     "example_2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
		},
		{
			name:     "example_3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		},
		{
			name: "long_reuse",
			s:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: []string{
				"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa",
			},
		},
	}

	implementations := []struct {
		name string
		fn   func(string, []string) bool
	}{
		{name: "wordBreak", fn: wordBreak},
		{name: "wordBreak2", fn: wordBreak2},
	}

	for _, impl := range implementations {
		for _, bm := range benchmarks {
			b.Run(impl.name+"/"+bm.name, func(b *testing.B) {
				var result bool
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					result = impl.fn(bm.s, bm.wordDict)
				}
				wordBreakBenchResult = result
			})
		}
	}
}
