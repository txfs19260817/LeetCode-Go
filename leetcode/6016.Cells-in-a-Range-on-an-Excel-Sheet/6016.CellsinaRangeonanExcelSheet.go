package leetcode

func cellsInRange(s string) []string {
	var ans []string
	minCol, minRow, maxCol, maxRow := s[0], s[1], s[3], s[4]
	for c := byte(0); c <= maxCol-minCol; c++ {
		for r := byte(0); r <= maxRow-minRow; r++ {
			ans = append(ans, string([]byte{minCol + c, minRow + r}))
		}
	}
	return ans
}
