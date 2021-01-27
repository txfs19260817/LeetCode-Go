package _209_Minimum_Size_Subarray_Sum

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, res, flag := 0, len(nums), false
	for l, r := 0, 0; l < len(nums); {
		if r < len(nums) && sum < s {
			sum += nums[r]
			r++
		} else {
			sum -= nums[l]
			l++
		}
		if r-l <= res && sum >= s {
			flag = true
			res = r - l
		}
	}
	if !flag {
		return 0
	}
	return res
}
