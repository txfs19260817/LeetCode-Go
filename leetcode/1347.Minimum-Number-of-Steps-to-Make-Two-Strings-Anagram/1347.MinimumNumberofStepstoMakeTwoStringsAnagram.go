package leetcode

func minSteps(s string, t string) int {
	ans, count := 0, map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}
	for _, c := range count {
		ans += abs(c)
	}
	return ans / 2
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
