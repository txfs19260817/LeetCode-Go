package leetcode

import (
	"math"
)

func shortestDistance(grid [][]int) int {
	ans, m, n := math.MaxInt, len(grid), len(grid[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	totalDist, spaceValue := make([][]int, m), 0
	for i := range totalDist {
		totalDist[i] = make([]int, n)
	}

	bfs := func(i, j int) {
		q := [][2]int{{i, j}}

		for l := 1; len(q) > 0; l++ {
			var p [][2]int
			for _, cur := range q {
				for _, dir := range dirs {
					ni, nj := cur[0]+dir[0], cur[1]+dir[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == spaceValue {
						totalDist[ni][nj] += l
						grid[ni][nj]--
						p = append(p, [2]int{ni, nj})
					}
				}
			}
			q = p
		}
		spaceValue--
		return
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				bfs(i, j)
			}
		}
	}

	for i, td := range totalDist {
		for j, v := range td {
			if grid[i][j] == spaceValue {
				ans = min(ans, v)
			}
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
