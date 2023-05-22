package leetcode

func productExceptSelf(nums []int) []int {
	ans, r := make([]int, len(nums)), 1
	ans[0] = 1
	// `ans` stores prefix products of `nums`,
	// but exclude the last because ans[n-1] = prod(num[0...i-2])
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	// use `r` to accumulate the product of elements on the right side of `i`
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}
