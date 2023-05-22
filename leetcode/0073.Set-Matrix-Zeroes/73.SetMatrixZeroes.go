package leetcode

func setZeroes(matrix [][]int) {
	var row0, col0 bool
	for _, l := range matrix {
		if l[0] == 0 {
			col0 = true
			break
		}
	}
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if col0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
	if row0 {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
}

func setZeroes1(matrix [][]int) {
	rows, cols := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	for i, l := range matrix {
		for j, v := range l {
			if v == 0 {
				rows[i], cols[j] = true, true
			}
		}
	}
	for i, l := range matrix {
		for j := range l {
			if rows[i] || cols[j] {
				l[j] = 0
			}
		}
	}
}
