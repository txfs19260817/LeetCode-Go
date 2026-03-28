package leetcode

import "strings"

func canChange(start string, target string) bool {
	if strings.Replace(start, "_", "", -1) != strings.Replace(target, "_", "", -1) {
		return false
	}
	var j int
	for i := range start {
		if start[i] != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && (start[i] == 'L' && i < j || start[i] == 'R' && i > j) {
				return false
			}
			j++
		}
	}
	return true
}
