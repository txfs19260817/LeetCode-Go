package _348_Design_Tic_Tac_Toe

type TicTacToe struct {
	rows, cols        []int
	diag, antiDiag, n int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{rows: make([]int, n), cols: make([]int, n), n: n}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	if player == 1 {
		this.rows[row]++
		if this.rows[row] == this.n {
			return player
		}
		this.cols[col]++
		if this.cols[col] == this.n {
			return player
		}
		if row == col {
			this.diag++
			if this.diag == this.n {
				return player
			}
		}
		if row+col == this.n-1 {
			this.antiDiag++
			if this.antiDiag == this.n {
				return player
			}
		}
	} else {
		this.rows[row]--
		if this.rows[row] == -this.n {
			return player
		}
		this.cols[col]--
		if this.cols[col] == -this.n {
			return player
		}
		if row == col {
			this.diag--
			if this.diag == -this.n {
				return player
			}
		}
		if row+col == this.n-1 {
			this.antiDiag--
			if this.antiDiag == -this.n {
				return player
			}
		}
	}
	return 0
}
