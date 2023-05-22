package leetcode

func rectangleArea(rectangles [][]int) int {
	ans, m := 0, int(1e9+7)
	var all [][]int
	for _, cur := range rectangles {
		breakdown(&all, cur, 0)
	}
	for _, rec := range all {
		ans += (rec[2] - rec[0]) * (rec[3] - rec[1])
		ans %= m
	}
	return ans
}

func breakdown(all *[][]int, cur []int, idx int) {
	if idx >= len(*all) {
		*all = append(*all, cur)
		return
	}

	rec := (*all)[idx]
	if !(cur[0] < rec[2] && cur[1] < rec[3] && cur[2] > rec[0] && cur[3] > rec[1]) {
		breakdown(all, cur, idx+1)
		return
	}
	if cur[0] < rec[0] {
		breakdown(all, []int{cur[0], cur[1], rec[0], cur[3]}, idx+1)
	}
	if cur[1] < rec[1] {
		breakdown(all, []int{max(rec[0], cur[0]), cur[1], min(rec[2], cur[2]), rec[1]}, idx+1)
	}
	if cur[2] > rec[2] {
		breakdown(all, []int{rec[2], cur[1], cur[2], cur[3]}, idx+1)
	}
	if cur[3] > rec[3] {
		breakdown(all, []int{max(rec[0], cur[0]), rec[3], min(rec[2], cur[2]), cur[3]}, idx+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
