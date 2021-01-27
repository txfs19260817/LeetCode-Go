package _080_Remove_Duplicates_from_Sorted_Array_II

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	var s, f, uniquePt, count int
	for ; f < len(nums); f++ {
		if nums[uniquePt] < nums[f] {
			uniquePt, count = f, 0
		}
		if count < 2 {
			nums[s] = nums[f]
			count++
			s++
		}
	}
	return s
}
