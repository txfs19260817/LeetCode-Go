package _930_Binary_Subarrays_With_Sum

func numSubarraysWithSum(A []int, S int) int {
	ans, sumi, m := 0, 0, map[int]int{0: 1}
	for _, num := range A {
		sumi += num
		sumj := sumi - S
		ans += m[sumj]
		m[sumi] += 1
	}
	return ans
}
