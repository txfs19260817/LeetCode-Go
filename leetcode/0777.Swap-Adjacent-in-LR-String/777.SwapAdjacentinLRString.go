package leetcode

import "strings"

func canTransform(start string, result string) bool {
	// L/R can't cross each other, so they should match
	if strings.Replace(start, "X", "", -1) != strings.Replace(result, "X", "", -1) {
		return false
	}
	j := 0
	for i, s := range start {
		if s != 'X' {
			for result[j] == 'X' {
				j++
			}
			// L can't move right; R can't move left
			if i != j && s == 'L' != (i > j) { // if i != j && (s == 'L' && i < j || s == 'R' && i > j) {
				return false
			}

			j++
		}
	}
	return true
}
