package leetcode

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
			return mid
		}
		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] < nums[r] {
		return r
	}
	return l
}
