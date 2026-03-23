package leetcode

func numberOfLines(widths []int, s string) []int {
	ans := []int{1, 100}
	for _, c := range s {
		w := widths[int(c-'a')]
		if ans[1] >= w {
			ans[1] -= w
		} else {
			ans[0]++
			ans[1] = 100 - w
		}
	}
	ans[1] = 100 - ans[1]
	return ans
}
