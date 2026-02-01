package leetcode

func minimumSubarrayLength(nums []int, k int) int {
	ans := len(nums) + 1
	for i, x := range nums {
		if x >= k { // single number
			return 1
		}
		for j := i - 1; j >= 0 && (nums[j]|x) != nums[j]; j-- {
			nums[j] |= x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}
