package _153_Find_Minimum_in_Rotated_Sorted_Array

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] >= nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
