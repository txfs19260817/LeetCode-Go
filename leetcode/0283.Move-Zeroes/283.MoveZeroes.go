package _283_Move_Zeroes

func moveZeroes(nums []int) {
	for s, f := 0, 0; f < len(nums); f++ {
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}
	}
}

func moveZeroes1(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i, count := len(nums)-1, 0; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums)-count; j++ {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			count++
		}
	}
}
