package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{binarySearch(nums, target, true), binarySearch(nums, target, false)}
}

func binarySearch(nums []int, target int, first bool) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else if nums[mid] < target {
			l = mid
		} else {
			if first {
				r = mid
			} else {
				l = mid
			}
		}
	}
	if !first {
		l, r = r, l
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
