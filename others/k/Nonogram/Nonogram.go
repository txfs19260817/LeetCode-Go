package k

import "slices"

func nonogram(mat, rows, cols [][]int) bool {
	if len(mat) != len(rows) || len(mat[0]) != len(cols) {
		return false
	}
	// rows
	for i := 0; i < len(mat); i++ {
		var cnt int
		row := make([]int, 0, len(rows[i]))
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				cnt++
			} else {
				if cnt > 0 {
					row = append(row, cnt)
				}
				cnt = 0
			}
		}
		if cnt > 0 {
			row = append(row, cnt)
		}
		if !slices.Equal(row, rows[i]) {
			return false
		}
	}
	// cols
	for j := 0; j < len(mat[0]); j++ {
		var cnt int
		col := make([]int, 0, len(cols[j]))
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 0 {
				cnt++
			} else {
				if cnt > 0 {
					col = append(col, cnt)
				}
				cnt = 0
			}
		}
		if cnt > 0 {
			col = append(col, cnt)
		}
		if !slices.Equal(col, cols[j]) {
			return false
		}
	}
	return true
}

func nonogram_recursion(mat, rows, cols [][]int) bool {
	// assume len(mat) == len(rows); len(mat[0]) == len(cols)
	m, n := len(mat), len(mat[0])
	dirs := [2][2]int{{0, 1}, {1, 0}}
	var dfs func(i, j int, dir [2]int, inst []int) bool
	dfs = func(i, j int, dir [2]int, inst []int) bool {
		if len(inst) == 0 {
			return true
		}
		firstInst := inst[0]
		inst = inst[1:]

		countConsecutive0 := 0
		ni, nj := i, j
		for ; ni < m && nj < n; ni, nj = ni+dir[0], nj+dir[1] {
			if mat[ni][nj] == 0 {
				countConsecutive0++
			} else if mat[ni][nj] == 1 && countConsecutive0 > 0 {
				break
			}
		}
		if countConsecutive0 == firstInst {
			return dfs(ni, nj, dir, inst)
		}
		return false
	}
	for i, row := range rows {
		if !dfs(i, 0, dirs[0], row) {
			return false
		}
	}
	for j, col := range cols {
		if !dfs(0, j, dirs[1], col) {
			return false
		}
	}
	return true
}
