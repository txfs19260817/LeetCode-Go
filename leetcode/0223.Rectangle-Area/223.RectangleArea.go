package _223_Rectangle_Area

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	s1, s2 := (ax2-ax1)*(ay2-ay1), (bx2-bx1)*(by2-by1)
	if ax1 >= bx2 || ax2 <= bx1 || ay1 >= by2 || ay2 <= by1 {
		return s1 + s2
	}
	return s1 + s2 - (min(ax2, bx2)-max(ax1, bx1))*(min(ay2, by2)-max(ay1, by1))
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
