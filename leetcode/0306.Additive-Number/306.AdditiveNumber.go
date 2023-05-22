package leetcode

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 0 {
			break
		}
		for secEnd := firstEnd + 1; max(firstEnd, secEnd-firstEnd) <= len(num)-secEnd; secEnd++ {
			if num[firstEnd+1] == '0' && secEnd > firstEnd+1 {
				break
			}
			x1, _ := strconv.Atoi(num[:firstEnd+1])
			x2, _ := strconv.Atoi(num[firstEnd+1 : secEnd+1])
			if dfs(num, x1, x2, secEnd+1) {
				return true
			}
		}
	}
	return false
}

func dfs(num string, x1, x2, remainStart int) bool {
	if remainStart == len(num) {
		return true
	}
	if sum, sumStr := x1+x2, strconv.Itoa(x1+x2); strings.HasPrefix(num[remainStart:], sumStr) {
		return dfs(num, x2, sum, remainStart+len(sumStr))
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
