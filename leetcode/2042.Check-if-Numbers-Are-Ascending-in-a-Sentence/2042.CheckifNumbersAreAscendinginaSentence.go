package leetcode

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	parts, num := strings.Split(s, " "), -1
	for _, part := range parts {
		a, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		if a <= num {
			return false
		}
		num = a
	}
	return true
}
