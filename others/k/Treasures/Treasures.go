package k

func findLegalMoves(board [][]int, i, j int) (ans [][]int) {
	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		r, c := i+dir[0], j+dir[1]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && board[r][c] == 0 {
			ans = append(ans, []int{r, c})
		}
	}
	return ans
}

// (endI, endJ) is the end position
func isAllPositionsPassed(board [][]int, endI, endJ int) bool {
	mat := board
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[0]) || mat[i][j] != 0 {
			return
		}
		mat[i][j] = -1
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	dfs(endI, endJ)
	for _, row := range mat {
		for _, v := range row {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

func findAllTreasures(board [][]int, start, end []int) (ans [][]int) {
	mat := board
	var treasures int
	for _, row := range mat {
		for _, v := range row {
			if v == 1 {
				treasures++
			}
		}
	}
	var dfs func(i, j, remained int, path [][]int)
	dfs = func(i, j, remained int, path [][]int) {
		if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[0]) || mat[i][j] == -1 || mat[i][j] == 2 {
			return
		}
		path = append(path, []int{i, j})
		tmp := mat[i][j]
		if tmp == 1 {
			remained--
		}
		if i == end[0] && j == end[1] && remained == 0 {
			if ans == nil || len(ans) > len(path) {
				ans = path
			}
			return
		}
		mat[i][j] = 2
		dfs(i+1, j, remained, path)
		dfs(i, j+1, remained, path)
		dfs(i-1, j, remained, path)
		dfs(i, j-1, remained, path)
		mat[i][j] = tmp
		path = path[:len(path)-1]
	}
	dfs(start[0], start[1], treasures, [][]int{})
	return
}
