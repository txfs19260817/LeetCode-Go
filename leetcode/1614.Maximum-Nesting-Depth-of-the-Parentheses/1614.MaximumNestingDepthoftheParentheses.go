package leetcode

func maxDepth(s string) int {
	var count, ans int
	for _, c := range s {
		if c == '(' {
			count++
			if count > ans {
				ans = count
			}
		} else if c == ')' {
			count--
		}
	}
	return ans
}
