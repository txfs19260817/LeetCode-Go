package leetcode

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	ans, newSp := 0, append([]int{bottom - 1, top + 1}, special...)
	sort.Ints(newSp)
	for i := 1; i < len(newSp); i++ {
		ans = max(ans, newSp[i]-newSp[i-1]-1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
