package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toBoard(rows []string) [][]byte {
	board := make([][]byte, len(rows))
	for i := range rows {
		board[i] = []byte(rows[i])
	}
	return board
}

func toRows(board [][]byte) []string {
	rows := make([]string, len(board))
	for i := range board {
		rows[i] = string(board[i])
	}
	return rows
}

func TestDropFigure(t *testing.T) {
	tests := []struct {
		name     string
		board    []string
		expected []string
	}{
		{
			name: "drop to bottom",
			board: []string{
				"....",
				".F..",
				".F..",
				"....",
			},
			expected: []string{
				"....",
				"....",
				".F..",
				".F..",
			},
		},
		{
			name: "stop by obstacle with minimal column gap",
			board: []string{
				"..F..",
				".FFF.",
				".....",
				"..#..",
				".....",
			},
			expected: []string{
				".....",
				"..F..",
				".FFF.",
				"..#..",
				".....",
			},
		},
		{
			name: "already blocked no movement",
			board: []string{
				".F.",
				".#.",
				"...",
			},
			expected: []string{
				".F.",
				".#.",
				"...",
			},
		},
		{
			name: "no figure",
			board: []string{
				".#.",
				"...",
			},
			expected: []string{
				".#.",
				"...",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DropFigure(toBoard(tc.board))
			assert.Equal(t, tc.expected, toRows(got))
		})
	}
}
