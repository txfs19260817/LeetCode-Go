package _330_Patching_Array

func minPatches(nums []int, n int) int {
	var ans int
	for i, patch := 0, 1; patch <= n; {
		if i < len(nums) && patch >= nums[i] {
			patch += nums[i]
			i++
		} else {
			patch += patch
			ans++
		}
	}
	return ans
}
