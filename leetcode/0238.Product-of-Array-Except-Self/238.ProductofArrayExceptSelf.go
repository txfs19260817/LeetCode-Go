package _238_Product_of_Array_Except_Self

func productExceptSelf(nums []int) []int {
	ans, r := make([]int, len(nums)), 1
	ans[0] = 1
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}
