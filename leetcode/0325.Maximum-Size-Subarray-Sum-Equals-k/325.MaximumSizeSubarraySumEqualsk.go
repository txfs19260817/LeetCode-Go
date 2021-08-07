package _325_Maximum_Size_Subarray_Sum_Equals_k

func maxSubArrayLen(nums []int, k int) int {
	ans, preSum, sum2idx := 0, 0, map[int]int{0: -1}
	for i, num := range nums {
		preSum += num
		if preIdx, ok := sum2idx[preSum-k]; ok {
			if ans < i-preIdx {
				ans = i - preIdx
			}
		}
		if _, ok := sum2idx[preSum]; !ok {
			sum2idx[preSum] = i
		}
	}
	return ans
}
