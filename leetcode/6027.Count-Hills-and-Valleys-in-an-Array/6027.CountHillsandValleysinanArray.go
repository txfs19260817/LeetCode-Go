package _027_Count_Hills_and_Valleys_in_an_Array

func countHillValley(nums []int) int {
	var ans int
	var last string
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		var l, r int
		for l = i + 1; l < len(nums)-1 && nums[l] == nums[i]; l++ {
		}
		for r = i - 1; r > 0 && nums[r] == nums[i]; r-- {
		}
		if nums[i] > nums[l] && nums[i] > nums[r] && last != "h" {
			ans++
			last = "h"
		} else if nums[i] < nums[l] && nums[i] < nums[r] && last != "v" {
			ans++
			last = "v"
		}
	}
	return ans
}
