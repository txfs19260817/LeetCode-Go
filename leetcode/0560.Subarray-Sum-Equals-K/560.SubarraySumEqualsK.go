package leetcode

func subarraySum(nums []int, k int) int {
	ans, preSum, m := 0, 0, map[int]int{0: 1}
	for _, num := range nums {
		preSum += num
		ans += m[preSum-k]
		m[preSum]++
	}
	return ans
}
