package leetcode

var (
	visit [][]int
	dirX  = []int{0, -1, 0, 1}
	dirY  = []int{-1, 0, 1, 0}
)

func largestIsland(grid [][]int) int {
	ans, m, n := 0, len(grid), len(grid[0])
	visit = make([][]int, m)
	for i := range visit {
		visit[i] = make([]int, n)
	}
	table, value := make([]int, m*n+2), 1
	for i, line := range grid {
		for j, e := range line {
			if e == 1 && visit[i][j] == 0 {
				table[value] = dfs(i, j, value, grid)
				value++
			}
		}
	}
	for i, line := range grid {
		for j, e := range line {
			if e != 0 {
				continue
			}
			sq, vmap := 1, map[int]bool{}
			if i-1 >= 0 && visit[i-1][j] > 0 {
				sq += table[visit[i-1][j]]
				vmap[visit[i-1][j]] = true
			}
			if j-1 >= 0 && visit[i][j-1] > 0 && !vmap[visit[i][j-1]] {
				sq += table[visit[i][j-1]]
				vmap[visit[i][j-1]] = true
			}
			if i+1 < m && visit[i+1][j] > 0 && !vmap[visit[i+1][j]] {
				sq += table[visit[i+1][j]]
				vmap[visit[i+1][j]] = true
			}
			if j+1 < n && visit[i][j+1] > 0 && !vmap[visit[i][j+1]] {
				sq += table[visit[i][j+1]]
				vmap[visit[i][j+1]] = true
			}
			if sq > ans {
				ans = sq
			}
		}
	}
	if ans == 0 {
		return m * n
	}
	return ans
}

func dfs(i, j, value int, grid [][]int) int {
	count := 1
	visit[i][j] = value
	for k := 0; k < 4; k++ {
		x, y := i+dirX[k], j+dirY[k]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && value != visit[x][y] && grid[x][y] == 1 {
			count += dfs(x, y, value, grid)
		}
	}
	return count
}
