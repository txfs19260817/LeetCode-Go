package _030_Matrix_Cells_in_Distance_Order

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDist := max(r0, R-1-r0) + max(c0, C-1-c0)
	bucket := make([][][]int, maxDist+1)
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			dist := abs(r-r0) + abs(c-c0)
			bucket[dist] = append(bucket[dist], []int{r, c})
		}
	}

	res := make([][]int, 0, R*C)
	for _, b := range bucket {
		res = append(res, b...)
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return abs(res[i][0]-r0)+abs(res[i][1]-c0) < abs(res[j][0]-r0)+abs(res[j][1]-c0)
	})

	return res
}
