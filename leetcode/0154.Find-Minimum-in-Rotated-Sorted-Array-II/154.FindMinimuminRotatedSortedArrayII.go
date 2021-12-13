package _154_Find_Minimum_in_Rotated_Sorted_Array_II

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}
