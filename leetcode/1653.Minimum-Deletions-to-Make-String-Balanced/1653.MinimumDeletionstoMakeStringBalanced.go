package leetcode

func minimumDeletions(s string) int {
	var dp, cntB int
	for _, c := range s {
		if c == 'b' { // reduce the problem to the string before it
			cntB++
		} else { // delete either this 'a' or all 'b's before
			dp = min(dp+1, cntB)
		}
	}
	return dp
}
