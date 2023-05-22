package leetcode

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	if len(s) < 4 || len(s) > 12 {
		return ans
	}
	dfs(&ans, s, []string{}, 3)
	return ans
}

func dfs(ans *[]string, s string, path []string, dots int) {
	if dots == 0 {
		if check(s) {
			path = append(path, s)
			*ans = append(*ans, strings.Join(path, "."))
		}
		return
	}
	for i := 1; i <= len(s) && i <= 3; i++ {
		if check(s[:i]) {
			path = append(path, s[:i])
			dfs(ans, s[i:], path, dots-1)
			path = path[:len(path)-1]
		}
	}
}

func check(s string) bool {
	if len(s) == 0 || s[0] == '0' && s != "0" {
		return false
	}
	if n, err := strconv.Atoi(s); err != nil || n > 255 || n < 0 {
		return false
	}
	return true
}
