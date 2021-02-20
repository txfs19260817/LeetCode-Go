package _136_Single_Number

func singleNumber(nums []int) int {
	ans := nums[0]
	for _, num := range nums[1:] {
		ans ^= num
	}
	return ans
}
