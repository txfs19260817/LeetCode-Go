package leetcode

func findJudge(n int, trust [][]int) int {
	degreeDelta := make([]int, n+1)
	for _, t := range trust {
		degreeDelta[t[0]]-- // out-degree+1
		degreeDelta[t[1]]++ // in-degree+1
	}
	for i := 1; i < len(degreeDelta); i++ {
		if degreeDelta[i] == n-1 {
			return i
		}
	}
	return -1
}
