package _219_Contains_Duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]bool{} // num:exist
	for i, num := range nums {
		if window[num] {
			return true
		}
		window[num] = true
		if i >= k {
			delete(window, nums[i-k])
		}
	}
	return false
}
