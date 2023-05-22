package leetcode

func isRectangleOverlap(rec1 []int, rec2 []int) bool { // [x1, y1, x2, y2]
	return rec1[0] < rec2[2] && rec1[1] < rec2[3] && rec1[2] > rec2[0] && rec1[3] > rec2[1]
}
