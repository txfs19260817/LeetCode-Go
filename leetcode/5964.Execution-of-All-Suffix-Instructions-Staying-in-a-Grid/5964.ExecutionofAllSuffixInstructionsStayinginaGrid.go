package leetcode

func executeInstructions(n int, startPos []int, s string) []int {
	dirs := map[byte][2]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	ans := make([]int, len(s))
	for i := range s {
		next := [2]int{startPos[0], startPos[1]}
		for j := i; j < len(s); j++ {
			d := dirs[s[j]]
			next[0] += d[0]
			next[1] += d[1]
			if next[0] >= n || next[0] < 0 || next[1] >= n || next[1] < 0 {
				break
			}
			ans[i]++
		}
	}
	return ans
}
