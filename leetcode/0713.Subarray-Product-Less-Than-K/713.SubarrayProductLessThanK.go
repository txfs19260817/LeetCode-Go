package _713_Subarray_Product_Less_Than_K

func numSubarrayProductLessThanK(nums []int, k int) int {
	prod, ans := 1, 0
	for l, r := 0, 0; l < len(nums); {
		if r < len(nums) && prod*nums[r] < k {
			prod *= nums[r]
			r++
		} else if l == r {
			l++
			r++
		} else {
			ans += r - l
			prod /= nums[l]
			l++
		}
	}
	return ans
}
