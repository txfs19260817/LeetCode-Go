package _542_01_Matrix

func updateMatrix(mat [][]int) [][]int {
	dirX, dirY, finished := []int{0, -1, 0, 1}, []int{-1, 0, 1, 0}, false
	duplicate, visited := make([][]int, len(mat)), make([][]bool, len(mat))
	for i := range mat {
		visited[i] = make([]bool, len(mat[i]))
		duplicate[i] = make([]int, len(mat[i]))
		copy(duplicate[i], mat[i])
	}
	for !finished {
		finished = true
		for i, line := range mat {
			for j, e := range line {
				if visited[i][j] {
					continue
				}
				if e == 0 {
					visited[i][j] = true
					continue
				}
				finished = false
				minAround := len(mat) * len(mat[0])
				for k := 0; k < 4; k++ {
					if x, y := i+dirX[k], j+dirY[k]; x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] < minAround {
						minAround = mat[x][y]
					}
				}
				if mat[i][j] <= minAround {
					duplicate[i][j] = minAround + 1
				} else {
					visited[i][j] = true
				}
			}
		}
		mat, duplicate = duplicate, mat
	}
	return mat
}
