package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantEmpty    bool
		wantExact    bool
		wantNonEmpty bool
	}{
		{
			name:      "leetcode_example_unique",
			args:      args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want:      "wertf",
			wantExact: true,
		},
		{
			name:         "simple_order",
			args:         args{words: []string{"z", "x"}},
			wantNonEmpty: true,
		},
		{
			name:      "prefix_invalid",
			args:      args{words: []string{"abc", "ab"}},
			wantEmpty: true,
		},
		{
			name:      "cycle_invalid",
			args:      args{words: []string{"z", "x", "z"}},
			wantEmpty: true,
		},
		{
			name:         "single_word",
			args:         args{words: []string{"abc"}},
			wantNonEmpty: true,
		},
		{
			name:         "repeated_words",
			args:         args{words: []string{"abc", "abc"}},
			wantNonEmpty: true,
		},
		{
			name:         "no_constraints_multiple_letters",
			args:         args{words: []string{"a", "b", "c"}},
			wantNonEmpty: true,
		},
		{
			name:         "classic_partial_order",
			args:         args{words: []string{"baa", "abcd", "abca", "cab", "cad"}},
			wantNonEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alienOrder(tt.args.words)

			if tt.wantEmpty {
				assert.Equal(t, "", got)
				return
			}

			if tt.wantExact {
				assert.Equal(t, tt.want, got)
				return
			}

			if tt.wantNonEmpty {
				assert.NotEmpty(t, got)
				assert.True(t, isPermutationOfPresent(tt.args.words, got),
					"got=%q is not a permutation of letters present in words=%v", got, tt.args.words)

				ok, reason := respectsOrdering(tt.args.words, got)
				assert.True(t, ok, "got=%q violates constraints: %s; words=%v", got, reason, tt.args.words)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

// isPermutationOfPresent checks output contains each present letter exactly once.
func isPermutationOfPresent(words []string, order string) bool {
	present := [26]bool{}
	presentCount := 0
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			c := int(w[i] - 'a')
			if 0 <= c && c < 26 && !present[c] {
				present[c] = true
				presentCount++
			}
		}
	}

	if len(order) != presentCount {
		return false
	}
	seen := [26]bool{}
	for i := 0; i < len(order); i++ {
		c := int(order[i] - 'a')
		if c < 0 || c >= 26 {
			return false
		}
		if !present[c] || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

// respectsOrdering validates the returned order respects constraints derived from adjacent words,
// and detects the invalid prefix case.
func respectsOrdering(words []string, order string) (bool, string) {
	pos := make([]int, 26)
	for i := range pos {
		pos[i] = -1
	}
	for i := 0; i < len(order); i++ {
		pos[int(order[i]-'a')] = i
	}

	for i := 0; i+1 < len(words); i++ {
		w1, w2 := words[i], words[i+1]

		// invalid prefix: "abc" before "ab"
		if len(w1) > len(w2) && strings.HasPrefix(w1, w2) {
			return false, "invalid prefix (longer word before its prefix)"
		}

		j := 0
		for j < len(w1) && j < len(w2) && w1[j] == w2[j] {
			j++
		}
		if j < len(w1) && j < len(w2) {
			u := int(w1[j] - 'a')
			v := int(w2[j] - 'a')
			if u < 0 || u >= 26 || v < 0 || v >= 26 {
				continue
			}
			if pos[u] == -1 || pos[v] == -1 {
				return false, "order missing letters"
			}
			if pos[u] > pos[v] {
				return false, "constraint violated: expected " + string(byte('a'+u)) + " before " + string(byte('a'+v))
			}
		}
	}
	return true, ""
}
