package leetcode

import "testing"

func Test_findNeedleInChunkedStrings(t *testing.T) {
	tests := []struct {
		name   string
		chunks []string
		needle string
		want   [2]int
	}{
		{
			name:   `chunks = ["a", "bcd", "ea", "de"], needle = "de"`,
			chunks: []string{"a", "bcd", "ea", "de"},
			needle: "de",
			want:   [2]int{1, 2},
		},
		{
			name:   `match inside a single chunk`,
			chunks: []string{"hello", "world"},
			needle: "llo",
			want:   [2]int{0, 2},
		},
		{
			name:   `match crosses chunk boundaries and skips empty chunks`,
			chunks: []string{"ab", "", "c", "de"},
			needle: "bcde",
			want:   [2]int{0, 1},
		},
		{
			name:   `return the earliest match in concatenation order`,
			chunks: []string{"ab", "cd", "abc"},
			needle: "abc",
			want:   [2]int{0, 0},
		},
		{
			name:   `match starts in a later non-empty chunk`,
			chunks: []string{"", "", "de"},
			needle: "de",
			want:   [2]int{2, 0},
		},
		{
			name:   `needle longer than remaining suffix`,
			chunks: []string{"ab", "c"},
			needle: "abcd",
			want:   [2]int{-1, -1},
		},
		{
			name:   `needle not found`,
			chunks: []string{"ab", "cd"},
			needle: "ef",
			want:   [2]int{-1, -1},
		},
		{
			name:   `empty needle returns the first valid position`,
			chunks: []string{"ab", "cd"},
			needle: "",
			want:   [2]int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNeedleInChunkedStrings(tt.chunks, tt.needle); got != tt.want {
				t.Errorf("findNeedleInChunkedStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
