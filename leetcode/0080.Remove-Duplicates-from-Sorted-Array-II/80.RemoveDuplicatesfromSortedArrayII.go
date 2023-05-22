package leetcode

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
	return process(2)
}

func removeDuplicates1(nums []int) int {
	var s, f, uniquePt, count int
	for ; f < len(nums); f++ {
		if nums[uniquePt] < nums[f] {
			uniquePt, count = f, 0
		}
		if count < 2 {
			nums[s] = nums[f]
			count++
			s++
		}
	}
	return s
}
