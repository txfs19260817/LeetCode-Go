package k

func nonogram(mat, rows, cols [][]int) bool {
	if len(mat) != len(rows) || len(mat[0]) != len(cols) {
		return false
	}
	// rows
	for i := 0; i < len(mat); i++ {
		var cnt int
		var row []int
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
		if !sliceEq(row, rows[i]) {
			return false
		}
	}
	// cols
	for j := 0; j < len(mat[0]); j++ {
		var cnt int
		var col []int
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
		if !sliceEq(col, cols[j]) {
			return false
		}
	}
	return true
}

func sliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
