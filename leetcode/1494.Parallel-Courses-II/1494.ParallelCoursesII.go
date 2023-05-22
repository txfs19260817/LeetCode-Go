package leetcode

import "math/bits"

// https://leetcode.com/problems/parallel-courses-ii/discuss/708263/Can-anyone-explain-the-bit-mask-method
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	pre, courseCnt := make([]int, n), 1<<n
	for _, relation := range relations {
		pre[relation[1]-1] |= 1 << (relation[0] - 1)
	}
	dp := make([]int, courseCnt)
	for i := 1; i < len(dp); i++ {
		dp[i] = n
	}
	for i := 0; i < courseCnt; i++ {
		var ex int // all courses that we can study now
		for j, p := range pre {
			if i&p == p {
				ex |= 1 << j
			}
		}
		ex &= ^i // remove what we have already studied
		for s := ex; s > 0; s = ex & (s - 1) {
			if bits.OnesCount(uint(s)) <= k {
				if v := dp[i] + 1; dp[i|s] > v {
					dp[i|s] = v // i|s: combine already studied courses with courses we can study at the current step
				}
			}
		}
	}
	return dp[len(dp)-1]
}
