package databricks

// TicTacToe represents a generalized n×m tic-tac-toe game
// where a player wins by getting k consecutive marks in a line.
type TicTacToe struct {
	n, m, k int
	board   map[[2]int]int // (row, col) -> player
	winner  int
}

// NewTicTacToe initializes an n×m board with win condition k in a row.
func NewTicTacToe(n, m, k int) *TicTacToe {
	return &TicTacToe{
		n:     n,
		m:     m,
		k:     k,
		board: make(map[[2]int]int),
	}
}

// directions: horizontal, vertical, main diagonal, anti-diagonal
var tttDirections = [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

// Move places a mark for player at (row, col).
// Returns 0 if no winner, or the winning player number (1 or 2).
func (t *TicTacToe) Move(row, col, player int) int {
	if t.winner != 0 {
		return t.winner
	}
	t.board[[2]int{row, col}] = player

	for _, d := range tttDirections {
		count := 1
		// Extend in the positive direction
		for i := 1; i < t.k; i++ {
			nr, nc := row+d[0]*i, col+d[1]*i
			if nr < 0 || nr >= t.n || nc < 0 || nc >= t.m {
				break
			}
			if t.board[[2]int{nr, nc}] != player {
				break
			}
			count++
		}
		// Extend in the negative direction
		for i := 1; i < t.k; i++ {
			nr, nc := row-d[0]*i, col-d[1]*i
			if nr < 0 || nr >= t.n || nc < 0 || nc >= t.m {
				break
			}
			if t.board[[2]int{nr, nc}] != player {
				break
			}
			count++
		}
		if count >= t.k {
			t.winner = player
			return player
		}
	}
	return 0
}
