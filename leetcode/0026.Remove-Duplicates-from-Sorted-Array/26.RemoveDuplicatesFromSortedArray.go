package _026_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	process := func(k int) int {
		var w int // writer pointer
		for _, v := range nums {
			if w < k || nums[w-k] != v {
				nums[w] = v
				w++
			}
		}
		return w
	}
	return process(1)
}

func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	s, f := 0, 1
	for s < len(nums)-1 {
		for nums[s] == nums[f] {
			f++
			if f == len(nums) {
				return s + 1
			}
		}
		s++
		nums[s] = nums[f]
	}
	return s + 1
}
