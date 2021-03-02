package _044_Wildcard_Matching

import "fmt"

func isMatch(s string, p string) bool {
	m := map[string]bool{} // i,j:bool
	var dp func(int, int) bool
	dp = func(i int, j int) bool {
		if j == len(p) {
			return i == len(s)
		}
		if i == len(s) {
			for ; j < len(p); j++ {
				if p[j] != '*' {
					return false
				}
			}
			return true
		}
		ans, key := false, fmt.Sprintf("%d,%d", i, j)
		if b, ok := m[key]; ok {
			return b
		}
		if s[i] == p[j] || p[j] == '?' {
			ans = dp(i+1, j+1)
		} else {
			if p[j] == '*' {
				ans = dp(i, j+1) || dp(i+1, j)
			} else {
				ans = false // 发现不匹配
			}
		}
		m[key] = ans
		return ans
	}
	return dp(0, 0)
}
