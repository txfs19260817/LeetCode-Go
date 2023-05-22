package leetcode

import "fmt"

func isMatch(s string, p string) bool {
	m := map[string]bool{}     // i,j:bool
	var dp func(int, int) bool // 指针不指向 *
	dp = func(i int, j int) bool {
		// 模式串 p 已经匹配完，如果 s 也恰好被匹配完，则说明成功匹配
		if j == len(p) {
			return i == len(s)
		}
		// e.g. s = "a", p = "ab*c*" => true
		if i == len(s) {
			// 如果能匹配空串，一定是字符串和 * 成对儿出现
			if (len(p)-j)%2 == 1 {
				return false
			}
			// 检查是否为 x*y*z* 这种形式
			for ; j+1 < len(p); j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		ans, key := false, fmt.Sprintf("%d,%d", i, j)
		if b, ok := m[key]; ok {
			return b
		}
		if s[i] == p[j] || p[j] == '.' {
			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i+1, j) || dp(i, j+2) // 多次或者0次
			} else {
				ans = dp(i+1, j+1) // 1�?
			}
		} else {
			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2) // 0 次
			} else {
				ans = false // 发现不匹配
			}
		}
		m[key] = ans
		return ans
	}
	return dp(0, 0)
}
