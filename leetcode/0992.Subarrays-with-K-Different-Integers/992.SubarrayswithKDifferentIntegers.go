package _992_Subarrays_with_K_Different_Integers

func subarraysWithKDistinct(A []int, K int) int {
	return slideWindow(A, K) - slideWindow(A, K-1)
}

// slideWindow returns the subarrays with AT MOST K Different Integers
func slideWindow(A []int, K int) int {
	ans, freq := 0, map[int]int{}
	for l, r := 0, 0; r < len(A); r++ {
		if freq[A[r]] == 0 {
			K--
		}
		freq[A[r]]++
		for K < 0 {
			freq[A[l]]--
			if freq[A[l]] == 0 {
				K++
			}
			l++
		}
		ans += r - l + 1
	}
	return ans
}
