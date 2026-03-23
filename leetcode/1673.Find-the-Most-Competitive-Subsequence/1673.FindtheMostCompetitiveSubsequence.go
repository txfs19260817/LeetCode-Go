package leetcode

func mostCompetitive(nums []int, k int) []int {
	ans := make([]int, 0, k)
	for i, x := range nums {
		for len(ans) > 0 && ans[len(ans)-1] > x && len(ans)+len(nums)-i > k {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, x)
		}
	}
	return ans
}
