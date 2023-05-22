package leetcode

func findErrorNums(nums []int) []int {
	dup, mismatch := -1, 1
	for _, n := range nums {
		if nums[abs(n)-1] < 0 {
			dup = abs(n)
		} else {
			nums[abs(n)-1] *= -1
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			mismatch = i + 1
			break
		}
	}
	return []int{dup, mismatch}
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
