package k

func effectiveMatrix(mat [][]int) bool {
	rows, cols := make([][]int, len(mat)), make([][]int, len(mat))
	for i := range rows {
		rows[i], cols[i] = make([]int, len(mat)), make([]int, len(mat))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			v := mat[i][j] - 1
			rows[i][v]++
			cols[j][v]++
			if rows[i][v] > 1 || cols[i][v] > 1 {
				return false
			}
		}
	}
	return true
}
