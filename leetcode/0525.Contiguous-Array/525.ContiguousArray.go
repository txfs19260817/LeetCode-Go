package _525_Contiguous_Array

func findMaxLength(nums []int) int {
	ans, count, cnt2idx := 0, 0, map[int]int{0: -1}
	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if preIdx, ok := cnt2idx[count]; ok {
			if ans < i-preIdx {
				ans = i - preIdx
			}
		} else {
			cnt2idx[count] = i
		}
	}
	return ans
}
