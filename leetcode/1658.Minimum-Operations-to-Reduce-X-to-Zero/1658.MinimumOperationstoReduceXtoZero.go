package leetcode

func minOperations(nums []int, x int) int {
	whole := 0
	for _, n := range nums {
		whole += n
	}
	target := whole - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	l, r, sum, res := 0, 0, 0, -1
	for l < len(nums) {
		if r < len(nums) && sum+nums[r] <= target {
			sum += nums[r]
			r++
		} else {
			sum -= nums[l]
			l++
		}
		if sum == target && r-l > res {
			res = r - l
		}
	}
	if res == -1 {
		return -1
	}
	return len(nums) - res
}
