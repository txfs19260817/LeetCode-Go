package am

func countSubarrays(arr []int, m int) int {
	ans, odd, prefix := 0, 0, make([]int, len(arr)+1) // prefix: how many subarray containing `odd` odd number
	for _, a := range arr {
		prefix[odd]++
		if a%2 == 1 {
			odd++
		}
		if odd >= m {
			ans += prefix[odd-m]
		}
	}
	return ans
}
