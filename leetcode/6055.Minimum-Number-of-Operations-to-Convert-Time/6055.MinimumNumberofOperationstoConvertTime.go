package leetcode

import (
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	ans, start, end := 0, time2int(current), time2int(correct)
	for start < end {
		if start+60 <= end {
			start += 60
		} else if start+15 <= end {
			start += 15
		} else if start+5 <= end {
			start += 5
		} else if start+1 <= end {
			start += 1
		}
		ans++
	}
	return ans
}

func time2int(t string) int {
	parts := strings.Split(t, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h*60 + m
}
