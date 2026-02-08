package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicTacToe_ColumnWin(t *testing.T) {
	// 4×6 board, k=4: player 1 wins with 4 in column 2
	game := NewTicTacToe(4, 6, 4)
	assert.Equal(t, 0, game.Move(0, 2, 1))
	assert.Equal(t, 0, game.Move(0, 3, 2))
	assert.Equal(t, 0, game.Move(1, 2, 1))
	assert.Equal(t, 0, game.Move(1, 3, 2))
	assert.Equal(t, 0, game.Move(2, 2, 1))
	assert.Equal(t, 0, game.Move(2, 3, 2))
	assert.Equal(t, 1, game.Move(3, 2, 1)) // player 1 wins
}

func TestTicTacToe_DiagonalWinPlayer2(t *testing.T) {
	// 4×4 board, k=3: player 2 wins with main diagonal
	game := NewTicTacToe(4, 4, 3)
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(0, 0, 2))
	assert.Equal(t, 0, game.Move(1, 0, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(3, 3, 1))
	assert.Equal(t, 2, game.Move(2, 2, 2)) // player 2 wins diagonal (0,0)-(1,1)-(2,2)
}

func TestTicTacToe_HorizontalWin(t *testing.T) {
	// 3×3 board, k=3: player 1 wins with top row horizontal
	game := NewTicTacToe(3, 3, 3)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(1, 0, 2))
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 1, game.Move(0, 2, 1)) // player 1 wins row 0
}

func TestTicTacToe_NoWinnerUntilLast(t *testing.T) {
	// 4×4 board, k=4: no winner until the final move completes the line
	game := NewTicTacToe(4, 4, 4)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(1, 0, 2))
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(0, 2, 1))
	assert.Equal(t, 0, game.Move(1, 2, 2))
	// Player 1 now has 3 in row 0, still not enough
	assert.Equal(t, 1, game.Move(0, 3, 1)) // player 1 wins with 4 in row 0
}

func TestTicTacToe_KEqualsOne(t *testing.T) {
	// k=1: first move always wins
	game := NewTicTacToe(3, 3, 1)
	assert.Equal(t, 1, game.Move(1, 1, 1)) // immediately wins
}

func TestTicTacToe_AntiDiagonalWin(t *testing.T) {
	// 5×5 board, k=3: player 2 wins with anti-diagonal
	game := NewTicTacToe(5, 5, 3)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(0, 2, 2))
	assert.Equal(t, 0, game.Move(3, 3, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(4, 4, 1))
	assert.Equal(t, 2, game.Move(2, 0, 2)) // player 2 wins anti-diagonal (0,2)-(1,1)-(2,0)
}

func TestTicTacToe_LargeBoard(t *testing.T) {
	// Large board 10000×10000, k=3: verify it doesn't allocate full grid
	game := NewTicTacToe(10000, 10000, 3)
	assert.Equal(t, 0, game.Move(5000, 5000, 1))
	assert.Equal(t, 0, game.Move(0, 0, 2))
	assert.Equal(t, 0, game.Move(5001, 5000, 1))
	assert.Equal(t, 0, game.Move(0, 1, 2))
	assert.Equal(t, 1, game.Move(5002, 5000, 1)) // player 1 wins vertical
}
