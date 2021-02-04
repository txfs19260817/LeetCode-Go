package _986_Interval_List_Intersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		if s := intersect(firstList[i], secondList[j]); s != nil {
			res = append(res, s)
		}
		if firstList[i][1] > secondList[j][1] {
			j++
		} else {
			i++
		}
	}
	return res
}

func intersect(a, b []int) []int {
	if a[1] < b[0] || b[1] < a[0] {
		return nil
	}
	return []int{max(a[0], b[0]), min(a[1], b[1])}
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
