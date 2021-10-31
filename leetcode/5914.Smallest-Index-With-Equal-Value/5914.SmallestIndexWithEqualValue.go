package _914_Smallest_Index_With_Equal_Value

func smallestEqual(nums []int) int {
	ans := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == i%10 && ans > i {
			ans = i
		}
	}
	if ans == len(nums) {
		ans = -1
	}
	return ans
}
