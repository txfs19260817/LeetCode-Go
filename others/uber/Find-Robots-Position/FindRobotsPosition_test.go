package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRobotsPosition(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		distance []int
		want     [][]int
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'O', 'E', 'E', 'E', 'X'},
				{'E', 'O', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
				{'X', 'E', 'O', 'E', 'E'},
				{'X', 'E', 'X', 'E', 'X'},
			},
			distance: []int{2, 2, 4, 1},
			want:     [][]int{{1, 1}},
		},
		{
			name: "Example 2",
			board: [][]byte{
				{'O', 'E', 'X', 'O', 'O'},
				{'E', 'O', 'X', 'O', 'X'},
				{'X', 'X', 'O', 'E', 'E'},
				{'E', 'O', 'E', 'O', 'E'},
				{'O', 'O', 'X', 'O', 'O'},
			},
			distance: []int{2, 1, 2, 4},
			want:     [][]int{{3, 1}},
		},
		{
			name: "Example 3",
			board: [][]byte{
				{'O', 'X', 'O'},
				{'E', 'O', 'X'},
				{'O', 'X', 'O'},
			},
			distance: []int{1, 1, 1, 1},
			want:     [][]int{{2, 2}, {0, 2}},
		},
		{
			name: "Example 4",
			board: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
			},
			distance: []int{1, 1, 1, 1},
			want:     [][]int{{2, 2}, {2, 0}, {1, 1}, {0, 2}, {0, 0}},
		},
		{
			name: "Example 5",
			board: [][]byte{
				{'O', 'E', 'E'},
				{'E', 'O', 'X'},
				{'E', 'X', 'O'},
			},
			distance: []int{3, 1, 1, 3},
			want:     nil, // Expecting empty slice, which is usually nil or empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, FindRobotsPosition(tt.board, tt.distance))
			assert.ElementsMatch(t, tt.want, FindRobotsPosition2(tt.board, tt.distance))
			assert.ElementsMatch(t, tt.want, FindRobotsPosition3(tt.board, tt.distance))
		})
	}
}

func setupLargeBoard() ([][]byte, []int) {
	// Create a 200x200 board
	rows, cols := 200, 200
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			board[i][j] = 'E'
		}
	}

	// Place a robot at (100, 100)
	board[100][100] = 'O'

	// Distance to boundaries:
	// Left: 100 - (-1) = 101
	// Top: 100 - (-1) = 101
	// Bottom: 200 - 100 = 100
	// Right: 200 - 100 = 100
	distance := []int{101, 101, 100, 100}
	return board, distance
}

func setupManyRobotsBoard() ([][]byte, []int) {
	// Create a 50x50 board full of robots
	rows, cols := 50, 50
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			board[i][j] = 'O'
		}
	}

	// Target distance for center (25, 25)
	// Left: 26
	// Top: 26
	// Bottom: 50 - 25 = 25
	// Right: 50 - 25 = 25
	distance := []int{26, 26, 25, 25}
	return board, distance
}

func TestFindRobotsPosition_Large(t *testing.T) {
	board, distance := setupLargeBoard()
	want := [][]int{{100, 100}}

	assert.ElementsMatch(t, want, FindRobotsPosition(board, distance))
	assert.ElementsMatch(t, want, FindRobotsPosition2(board, distance))
	assert.ElementsMatch(t, want, FindRobotsPosition3(board, distance))
}

func TestFindRobotsPosition_Large_ManyRobots(t *testing.T) {
	board, distance := setupManyRobotsBoard()
	want := [][]int{{25, 25}}

	assert.ElementsMatch(t, want, FindRobotsPosition(board, distance))
	assert.ElementsMatch(t, want, FindRobotsPosition2(board, distance))
	assert.ElementsMatch(t, want, FindRobotsPosition3(board, distance))
}

func BenchmarkFindRobotsPosition_Large(b *testing.B) {
	board, distance := setupLargeBoard()

	for b.Loop() {
		FindRobotsPosition(board, distance)
	}
}

func BenchmarkFindRobotsPosition2_Large(b *testing.B) {
	board, distance := setupLargeBoard()

	for b.Loop() {
		FindRobotsPosition2(board, distance)
	}
}

func BenchmarkFindRobotsPosition3_Large(b *testing.B) {
	board, distance := setupLargeBoard()

	for b.Loop() {
		FindRobotsPosition3(board, distance)
	}
}

func BenchmarkFindRobotsPosition_Large_ManyRobots(b *testing.B) {
	board, distance := setupManyRobotsBoard()

	for b.Loop() {
		FindRobotsPosition(board, distance)
	}
}

func BenchmarkFindRobotsPosition2_Large_ManyRobots(b *testing.B) {
	board, distance := setupManyRobotsBoard()

	for b.Loop() {
		FindRobotsPosition2(board, distance)
	}
}

func BenchmarkFindRobotsPosition3_Large_ManyRobots(b *testing.B) {
	board, distance := setupManyRobotsBoard()

	for b.Loop() {
		FindRobotsPosition3(board, distance)
	}
}
