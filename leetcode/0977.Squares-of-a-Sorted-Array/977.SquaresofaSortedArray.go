package leetcode

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for l, r, i := 0, len(nums)-1, len(ans)-1; l <= r; i-- {
		if sl, sr := nums[l]*nums[l], nums[r]*nums[r]; sl > sr {
			ans[i] = sl
			l++
		} else {
			ans[i] = sr
			r--
		}
	}
	return ans
}
