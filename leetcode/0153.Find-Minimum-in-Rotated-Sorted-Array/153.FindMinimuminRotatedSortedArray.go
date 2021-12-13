package _153_Find_Minimum_in_Rotated_Sorted_Array

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] >= nums[r] {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}

func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] > nums[r] {
		return nums[r]
	}
	return nums[0]
}
