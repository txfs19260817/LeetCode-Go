package _525_Contiguous_Array

func findMaxLength(nums []int) int {
	ans, count, m := 0, 0, map[int]int{} // count:index
	m[0] = -1                            // initialize index = -1 where count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if preIdx, ok := m[count]; ok {
			if i-preIdx > ans {
				ans = i - preIdx
			}
		} else {
			m[count] = i
		}
	}
	return ans
}
