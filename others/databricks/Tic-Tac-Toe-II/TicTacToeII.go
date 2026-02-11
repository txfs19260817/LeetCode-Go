package databricks

// ---------------------------------------------------------------------------
// Solution 1: O(k) per move — scan up to k cells in each direction
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// Solution 2: O(1) per move — endpoint run-length counters
// ---------------------------------------------------------------------------
//
// Key idea: for each of 4 direction pairs (8 half-directions), maintain
//
//     run[(row, col, halfDir)] = length of the maximal same-player
//                                 consecutive run starting at (row,col)
//                                 and extending in halfDir
//
// Only the two ENDPOINTS of each run need correct values, because future
// placements can only land on empty cells (i.e. outside every existing run).
//
// When placing at (r,c), for each direction pair (d, d'):
//   a = run at the backward neighbor going further backward
//   b = run at the forward  neighbor going further forward
//   total = a + 1 + b   (merged run through the new cell)
//   → update far-back endpoint's d-value  = total
//   → update far-forward endpoint's d'-value = total
//
// 4 direction pairs × O(1) each = O(1) total per move.

// TicTacToeO1 is the O(1)-per-move variant.
type TicTacToeO1 struct {
	n, m, k int
	board   map[[2]int]int // (row, col) -> player
	run     map[[3]int]int // (row, col, halfDirIdx) -> run length
	winner  int
}

// 8 half-directions, grouped in opposite pairs (even/odd).
var halfDirs = [8][2]int{
	{0, 1}, {0, -1}, // pair 0: right  / left
	{1, 0}, {-1, 0}, // pair 1: down   / up
	{1, 1}, {-1, -1}, // pair 2: ↘      / ↖
	{1, -1}, {-1, 1}, // pair 3: ↙      / ↗
}

func NewTicTacToeO1(n, m, k int) *TicTacToeO1 {
	return &TicTacToeO1{
		n: n, m: m, k: k,
		board: make(map[[2]int]int),
		run:   make(map[[3]int]int),
	}
}

func (t *TicTacToeO1) Move(row, col, player int) int {
	if t.winner != 0 {
		return t.winner
	}
	t.board[[2]int{row, col}] = player

	for pair := 0; pair < 4; pair++ {
		d, dOpp := pair*2, pair*2+1
		dx, dy := halfDirs[d][0], halfDirs[d][1]

		// a = backward run (neighbor in the –d direction, extending further back)
		pr, pc := row-dx, col-dy
		a := 0
		if t.board[[2]int{pr, pc}] == player {
			a = t.run[[3]int{pr, pc, dOpp}]
		}

		// b = forward run (neighbor in the +d direction, extending further forward)
		sr, sc := row+dx, col+dy
		b := 0
		if t.board[[2]int{sr, sc}] == player {
			b = t.run[[3]int{sr, sc, d}]
		}

		total := a + 1 + b

		// Update far endpoints of the merged run
		t.run[[3]int{row - dx*a, col - dy*a, d}] = total
		t.run[[3]int{row + dx*b, col + dy*b, dOpp}] = total

		if total >= t.k {
			t.winner = player
			return player
		}
	}
	return 0
}
