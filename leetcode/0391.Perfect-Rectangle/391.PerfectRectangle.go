package _391_Perfect_Rectangle

func isRectangleCover(rectangles [][]int) bool {
	point2cnt := map[[2]int]int{}
	area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	for _, rectangle := range rectangles {
		minX = min(minX, rectangle[0])
		minY = min(minY, rectangle[1])
		maxX = max(maxX, rectangle[2])
		maxY = max(maxY, rectangle[3])
		area += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
		point2cnt[[2]int{rectangle[0], rectangle[1]}]++
		point2cnt[[2]int{rectangle[2], rectangle[3]}]++
		point2cnt[[2]int{rectangle[0], rectangle[3]}]++
		point2cnt[[2]int{rectangle[2], rectangle[1]}]++
	}
	if area != (maxX-minX)*(maxY-minY) || point2cnt[[2]int{minX, minY}] != 1 || point2cnt[[2]int{maxX, maxY}] != 1 || point2cnt[[2]int{minX, maxY}] != 1 || point2cnt[[2]int{maxX, minY}] != 1 {
		return false
	}
	delete(point2cnt, [2]int{minX, minY})
	delete(point2cnt, [2]int{maxX, maxY})
	delete(point2cnt, [2]int{maxX, minY})
	delete(point2cnt, [2]int{minX, maxY})
	for _, cnt := range point2cnt {
		if cnt != 2 && cnt != 4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
