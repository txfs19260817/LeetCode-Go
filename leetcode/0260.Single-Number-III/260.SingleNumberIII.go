package _260_Single_Number_III

func singleNumber(nums []int) []int {
	xors, div, a, b := nums[0], 1, 0, 0
	for _, n := range nums[1:] {
		xors ^= n
	}
	for xors&div == 0 {
		div <<= 1 // find the first 1
	}
	for _, n := range nums {
		if div&n == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
