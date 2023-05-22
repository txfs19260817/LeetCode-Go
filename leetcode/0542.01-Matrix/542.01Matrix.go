package leetcode

// DP
func updateMatrix(mat [][]int) [][]int {
	dist := make([][]int, len(mat))
	for i := range dist {
		dist[i] = make([]int, len(mat[0]))
		for j := range dist[i] {
			if mat[i][j] != 0 {
				dist[i][j] = len(mat) * len(mat[0])
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if i > 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j > 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if i < len(mat)-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < len(mat[0])-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// BFS
func updateMatrix2(mat [][]int) [][]int {
	dirX, dirY, finished := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}, false
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
		mat = duplicate
	}
	return mat
}
