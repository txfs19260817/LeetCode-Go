package leetcode

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		s, f, symbol := i, getNextIdx(nums, i), nums[i]
		for symbol*nums[f] > 0 && symbol*nums[getNextIdx(nums, f)] > 0 {
			if s == f {
				if s == getNextIdx(nums, s) {
					break
				}
				return true
			}
			s, f = getNextIdx(nums, s), getNextIdx(nums, getNextIdx(nums, f))
		}

		s = i
		for symbol*nums[s] > 0 {
			next := getNextIdx(nums, s)
			nums[s] = 0
			s = next
		}
	}
	return false
}

func getNextIdx(nums []int, i int) int {
	return ((nums[i]+i)%len(nums) + len(nums)) % len(nums)
}
