package leetcode

func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] {
			start++
		} else if nums[start] < nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[end] >= target && target > nums[mid] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}
