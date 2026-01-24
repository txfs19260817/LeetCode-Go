package uber

// FindRobotsPosition returns the indices of all robots that match the given distance criteria.
// board is a m x n grid. distance is [left, top, bottom, right].
func FindRobotsPosition(board [][]byte, distance []int) [][]int {
	dirs := [4][2]int{
		{0, -1}, // Left
		{-1, 0}, // Top
		{1, 0},  // Bottom
		{0, 1},  // Right
	}
	m, n := len(board), len(board[0])
	var ans [][]int

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' {
				match := true
				for i, d := range dirs {
					dist := 0
					nr, nc := r, c

					// Calculate distance to nearest blocker
					for {
						nr += d[0]
						nc += d[1]
						dist++
						// Check boundary or blocker
						if nr < 0 || nr >= m || nc < 0 || nc >= n || board[nr][nc] == 'X' {
							break
						}
					}

					if dist != distance[i] {
						match = false
						break
					}
				}
				if match {
					ans = append(ans, []int{r, c})
				}
			}
		}
	}
	return ans
}

func FindRobotsPosition2(board [][]byte, distance []int) [][]int {
	m, n := len(board), len(board[0])
	left, up, down, right := make([][]int, m), make([][]int, m), make([][]int, m), make([][]int, m)
	for i := range m {
		left[i], up[i], down[i], right[i] = make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	}

	// row-wise: l->r
	for i := range m {
		prev := -1
		for j := range n {
			if board[i][j] == 'X' {
				prev = j
			} else {
				left[i][j] = j - prev
			}
		}

		next := n
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' {
				next = j
			} else {
				right[i][j] = next - j
			}
		}
	}

	// col-wise: u->d
	for j := range n {
		prev := -1
		for i := range n {
			if board[i][j] == 'X' {
				prev = i
			} else {
				up[i][j] = i - prev
			}
		}

		next := n
		for i := n - 1; i >= 0; i-- {
			if board[i][j] == 'X' {
				next = i
			} else {
				down[i][j] = next - i
			}
		}
	}

	var ans [][]int
	for i, row := range board {
		for j, v := range row {
			if v == 'O' && left[i][j] == distance[0] && up[i][j] == distance[1] && down[i][j] == distance[2] && right[i][j] == distance[3] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func FindRobotsPosition3(board [][]byte, distance []int) [][]int {
	m, n := len(board), len(board[0])
	L, T, B, R := distance[0], distance[1], distance[2], distance[3]

	// Pass 1: top-down compute up + per-row left/right, filter candidates by (L,T,R)
	lastBlockerRow := make([]int, n)
	for j := range n {
		lastBlockerRow[j] = -1 // 上边界视为 blocker
	}

	candSet := map[int]bool{}
	rightRow := make([]int, n) // 只复用这一行的临时数组

	for i := range m {
		// 先算这一行的 right distances（从右到左）
		nextBlockerCol := n // 右边界视为 blocker
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' {
				nextBlockerCol = j
				rightRow[j] = 0 // blocker 本身无意义
				continue
			}
			rightRow[j] = nextBlockerCol - j
		}

		// 再从左到右：算 left & up，并结合 right 过滤
		prevBlockerCol := -1 // 左边界视为 blocker
		for j := range n {
			if board[i][j] == 'X' {
				prevBlockerCol = j
				lastBlockerRow[j] = i
				continue
			}

			left := j - prevBlockerCol
			up := i - lastBlockerRow[j]
			right := rightRow[j]

			if board[i][j] == 'O' && left == L && up == T && right == R {
				candSet[i*n+j] = true
			}
		}
	}

	// Pass 2: bottom-up compute down, intersect with candidates
	nextBlockerRow := make([]int, n)
	for j := range nextBlockerRow {
		nextBlockerRow[j] = m // 下边界视为 blocker
	}

	var ans [][]int
	for i := m - 1; i >= 0; i-- {
		for j := range n {
			if board[i][j] == 'X' {
				nextBlockerRow[j] = i
				continue
			}

			if board[i][j] == 'O' && candSet[i*n+j] && nextBlockerRow[j]-i == B {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
