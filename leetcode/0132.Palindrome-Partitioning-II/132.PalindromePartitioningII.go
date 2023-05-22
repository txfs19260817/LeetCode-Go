package leetcode

func minCut(s string) int {
	g := make([][]bool, len(s))
	for i := range g {
		g[i] = make([]bool, len(s))
		for j := range g[i] {
			g[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1] // s[i+1..j-1] is a palindrome
		}
	}
	f := make([]int, len(s)) // f[i] = min{f[j]}+1 && s[j+1..i] is a palindrome
	for i := range f {
		if g[0][i] { // If s[0..i] is a palindrome, it needs 0 cut.
			continue
		}
		f[i] = 1 << 31
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] { // s[j+1..i] is a palindrome
				f[i] = f[j] + 1
			}
		}
	}
	return f[len(s)-1]
}
