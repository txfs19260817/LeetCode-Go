package _007_Reverse_Integer

func reverse(x int) int {
	var ans int
	for ; x != 0; x /= 10 {
		ans = ans*10 + x%10
	}
	if ans > 1<<31-1 || ans < -(1<<31) {
		return 0
	}
	return ans
}
