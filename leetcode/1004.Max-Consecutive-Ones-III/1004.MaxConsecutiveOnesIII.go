package _004_Max_Consecutive_Ones_III

func longestOnes(A []int, K int) int {
	var ans int
	for l, r := 0, 0; l < len(A); {
		if r < len(A) && ((A[r] == 0 && K > 0) || A[r] == 1) { // we can continue to add the next 1 even if the window has already contained K 0s.
			if A[r] == 0 {
				K--
			}
			r++
		} else {
			if A[l] == 0 {
				K++
			}
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	return ans
}
