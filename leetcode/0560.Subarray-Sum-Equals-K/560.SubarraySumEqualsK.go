package _560_Subarray_Sum_Equals_K

func subarraySum(nums []int, k int) int {
	ans, sumi, m := 0, 0, map[int]int{0: 1}
	for _, num := range nums {
		sumi += num
		sumj := sumi - k
		ans += m[sumj]
		m[sumi] += 1
	}
	return ans
}
