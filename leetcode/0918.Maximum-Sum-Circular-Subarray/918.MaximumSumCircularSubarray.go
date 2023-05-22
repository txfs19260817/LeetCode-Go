package leetcode

func maxSubarraySumCircular(A []int) int {
	curMax, curMin, ansMax, ansMin, sum := 0, 0, A[0], A[0], 0
	for _, n := range A {
		// Sum
		sum += n
		// max sub
		curMax = max(n, curMax+n)
		ansMax = max(ansMax, curMax)
		// min sub
		curMin = min(n, curMin+n)
		ansMin = min(ansMin, curMin)
	}
	// all elements < 0
	if ansMax < 0 {
		return ansMax
	}
	return max(ansMax, sum-ansMin)
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
