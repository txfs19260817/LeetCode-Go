package leetcode

import (
	"sort"
	"testing"
)

func assertSentencesEqual(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(got)
	sort.Strings(want)
	if len(got) != len(want) {
		t.Fatalf("wordBreak() = %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("wordBreak() = %v, want %v", got, want)
		}
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
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
		},
		{
			name:     "example_2",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
		},
		{
			name:     "example_3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		},
		{
			name:     "dense_paths",
			s:        "aaaaaaa",
			wordDict: []string{"a", "aa", "aaa", "aaaa"},
		},
	}

	implementations := []struct {
		name string
		fn   func(string, []string) []string
	}{
		{name: "wordBreak", fn: wordBreak},
		{name: "wordBreak2", fn: wordBreak2},
	}

	for _, impl := range implementations {
		for _, bm := range benchmarks {
			b.Run(impl.name+"/"+bm.name, func(b *testing.B) {
				b.ResetTimer()
				for b.Loop() {
					_ = impl.fn(bm.s, bm.wordDict)
				}
			})
		}
	}
}

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cat sand dog", "cats and dog"},
		},
		{
			name: "example 2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			name: "example 3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertSentencesEqual(t, wordBreak(tt.args.s, tt.args.wordDict), tt.want)
			assertSentencesEqual(t, wordBreak2(tt.args.s, tt.args.wordDict), tt.want)
		})
	}
}
