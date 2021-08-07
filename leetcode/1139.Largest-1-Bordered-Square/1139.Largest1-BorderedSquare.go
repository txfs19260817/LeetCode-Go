package _139_Largest_1_Bordered_Square

func largest1BorderedSquare(grid [][]int) int {
	ans, px, py := 0, make([][]int, len(grid)+1), make([][]int, len(grid)+1) // prefix-sum
	for i := range px {
		px[i], py[i] = make([]int, len(grid[0])+1), make([]int, len(grid[0])+1)
	}
	for i := 1; i < len(px); i++ {
		for j := 1; j < len(px[0]); j++ {
			px[i][j], py[i][j] = px[i][j-1]+grid[i-1][j-1], py[i-1][j]+grid[i-1][j-1]
		}
	}
	for i := 1; i < len(px); i++ {
		for j := 1; j < len(px[0]); j++ {
			for k := len(px[0]) - 1; k >= j; k-- {
				l := k - j + 1
				if l <= ans {
					break
				}
				if i+l-1 >= len(px) {
					continue
				}
				if px[i][k]-px[i][j-1] != l {
					continue
				}
				if px[i+l-1][k]-px[i+l-1][j-1] != l {
					continue
				}
				if py[i+l-1][j]-py[i-1][j] != l {
					continue
				}
				if py[i+l-1][k]-py[i-1][k] != l {
					continue
				}
				ans = l
			}
		}
	}
	return ans * ans
}
