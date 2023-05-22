package leetcode

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

func moveZeroes2(nums []int) {
	var l, r int
	for l, r = 0, 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
}
