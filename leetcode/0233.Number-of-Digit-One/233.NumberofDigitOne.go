package _233_Number_of_Digit_One

func countDigitOne(n int) int {
	// mulk 表示 10^k
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k
	var ans int
	for k, mulk := 0, 1; mulk <= n; k, mulk = k+1, mulk*10 {
		ans += (n/(10*mulk))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
