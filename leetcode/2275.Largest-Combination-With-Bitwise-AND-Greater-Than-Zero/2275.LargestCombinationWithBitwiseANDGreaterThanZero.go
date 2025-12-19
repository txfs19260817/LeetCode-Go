package leetcode

func largestCombination(candidates []int) int {
	var ans int
	for k := 0; k < 24; k++ {
		var cnt int
		for _, c := range candidates {
			if (c>>k)&1 == 1 {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return ans
}
