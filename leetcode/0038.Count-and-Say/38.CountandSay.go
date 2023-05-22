package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return count(countAndSay(n - 1))
}

func count(s string) string {
	ans := make([]string, 0, len(s)*2)
	var j int
	for i := 0; i < len(s); i = j {
		var cnt int
		for j = i; j < len(s) && s[i] == s[j]; j++ {
			cnt++
		}
		ans = append(ans, []string{strconv.Itoa(cnt), string(s[i])}...)
	}
	return strings.Join(ans, "")
}
