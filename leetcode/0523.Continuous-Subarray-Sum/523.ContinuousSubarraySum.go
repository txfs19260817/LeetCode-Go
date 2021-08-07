package _523_Continuous_Subarray_Sum

func checkSubarraySum(nums []int, k int) bool {
	rem2idx, remainder := map[int]int{0: -1}, 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if preIdx, ok := rem2idx[remainder]; ok {
			if i-preIdx > 1 {
				return true
			}
		} else {
			rem2idx[remainder] = i
		}
	}
	return false
}
