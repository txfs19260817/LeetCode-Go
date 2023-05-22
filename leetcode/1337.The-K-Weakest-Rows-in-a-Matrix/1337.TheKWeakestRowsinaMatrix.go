package leetcode

import (
	"sort"
)

type pair struct{ row, num int }

func kWeakestRows(mat [][]int, k int) []int {
	var ans []int
	n := len(mat[0])
	soliderNum := func(r int) int {
		if mat[r][0] == 0 {
			return 0
		}
		if mat[r][n-1] == 1 {
			return n
		}
		return 1 + sort.Search(n, func(i int) bool { return mat[r][i] == 0 || i < n-1 && mat[r][i] == 1 && mat[r][i+1] == 0 })
	}
	h := make([]pair, 0, len(mat))
	for r := 0; r < len(mat); r++ {
		h = append(h, pair{row: r, num: soliderNum(r)})
	}
	sort.Slice(h, func(i, j int) bool { return h[i].num < h[j].num || h[i].num == h[j].num && h[i].row < h[j].row })
	for i := 0; i < k; i++ {
		ans = append(ans, h[i].row)
	}
	return ans
}
