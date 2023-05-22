package leetcode

import "math"

// DP O(N√N)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		minV := 1 << 31
		for j := 1; j*j <= i; j++ {
			if v := dp[i-j*j]; minV > v {
				minV = v
			}
		}
		dp[i] = minV + 1
	}
	return dp[n]
}

// DP O(N^2)
func numSquares1(n int) int {
	dp := make([]int, n+1)
	for i := 1; float64(i) <= math.Sqrt(float64(n)); i++ {
		dp[i*i] = 1
	}
	for i := 2; i < len(dp); i++ {
		if dp[i] == 1 {
			continue
		}
		dp[i] = 1 << 31
		for l, r := 1, i-1; l <= r; l, r = l+1, r-1 {
			if v := dp[l] + dp[r]; dp[i] > v {
				dp[i] = v
			}
		}
	}
	return dp[n]
}

// Math
func numSquares2(n int) int {
	isSquare := func(a int) bool {
		b := int(math.Sqrt(float64(a)))
		return a == b*b
	}
	check4 := func(a int) bool { // n=4^a*(8m+7) can be formulated by the sum of square of 4 numbers
		for a%4 == 0 {
			a /= 4
		}
		return a%8 == 7
	}
	if isSquare(n) {
		return 1
	}
	if check4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}
