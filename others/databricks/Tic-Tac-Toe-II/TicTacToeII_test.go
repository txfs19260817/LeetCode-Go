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

// ---------------------------------------------------------------------------
// O(1) solution tests (TicTacToeO1) — same scenarios as above
// ---------------------------------------------------------------------------

func TestTicTacToeO1_ColumnWin(t *testing.T) {
	game := NewTicTacToeO1(4, 6, 4)
	assert.Equal(t, 0, game.Move(0, 2, 1))
	assert.Equal(t, 0, game.Move(0, 3, 2))
	assert.Equal(t, 0, game.Move(1, 2, 1))
	assert.Equal(t, 0, game.Move(1, 3, 2))
	assert.Equal(t, 0, game.Move(2, 2, 1))
	assert.Equal(t, 0, game.Move(2, 3, 2))
	assert.Equal(t, 1, game.Move(3, 2, 1))
}

func TestTicTacToeO1_DiagonalWinPlayer2(t *testing.T) {
	game := NewTicTacToeO1(4, 4, 3)
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(0, 0, 2))
	assert.Equal(t, 0, game.Move(1, 0, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(3, 3, 1))
	assert.Equal(t, 2, game.Move(2, 2, 2))
}

func TestTicTacToeO1_HorizontalWin(t *testing.T) {
	game := NewTicTacToeO1(3, 3, 3)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(1, 0, 2))
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 1, game.Move(0, 2, 1))
}

func TestTicTacToeO1_NoWinnerUntilLast(t *testing.T) {
	game := NewTicTacToeO1(4, 4, 4)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(1, 0, 2))
	assert.Equal(t, 0, game.Move(0, 1, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(0, 2, 1))
	assert.Equal(t, 0, game.Move(1, 2, 2))
	assert.Equal(t, 1, game.Move(0, 3, 1))
}

func TestTicTacToeO1_KEqualsOne(t *testing.T) {
	game := NewTicTacToeO1(3, 3, 1)
	assert.Equal(t, 1, game.Move(1, 1, 1))
}

func TestTicTacToeO1_AntiDiagonalWin(t *testing.T) {
	game := NewTicTacToeO1(5, 5, 3)
	assert.Equal(t, 0, game.Move(0, 0, 1))
	assert.Equal(t, 0, game.Move(0, 2, 2))
	assert.Equal(t, 0, game.Move(3, 3, 1))
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(4, 4, 1))
	assert.Equal(t, 2, game.Move(2, 0, 2))
}

func TestTicTacToeO1_LargeBoard(t *testing.T) {
	game := NewTicTacToeO1(10000, 10000, 3)
	assert.Equal(t, 0, game.Move(5000, 5000, 1))
	assert.Equal(t, 0, game.Move(0, 0, 2))
	assert.Equal(t, 0, game.Move(5001, 5000, 1))
	assert.Equal(t, 0, game.Move(0, 1, 2))
	assert.Equal(t, 1, game.Move(5002, 5000, 1))
}

func TestTicTacToeO1_OutOfOrderPlacement(t *testing.T) {
	// Place marks out of order to verify endpoint merging:
	// row 0: place at cols 0, 3, 1, 2 → should detect 4-in-a-row on last move
	game := NewTicTacToeO1(5, 5, 4)
	assert.Equal(t, 0, game.Move(0, 0, 1)) // [X . . . .]
	assert.Equal(t, 0, game.Move(1, 0, 2))
	assert.Equal(t, 0, game.Move(0, 3, 1)) // [X . . X .]
	assert.Equal(t, 0, game.Move(1, 1, 2))
	assert.Equal(t, 0, game.Move(0, 1, 1)) // [X X . X .]
	assert.Equal(t, 0, game.Move(1, 2, 2))
	assert.Equal(t, 1, game.Move(0, 2, 1)) // [X X X X .] → bridges the gap, 4-in-a-row
}
