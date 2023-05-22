package leetcode

import "strings"

func isValidSerialization(preorder string) bool {
	nodes, degree := strings.Split(preorder, ","), 1 // degree is init with 1 since root has no in-degree
	for _, node := range nodes {
		degree--
		if degree < 0 { // degree should not be non-pos number since we have not visited its children yet
			return false
		}
		if node != "#" {
			degree += 2
		}
	}
	return degree == 0
}
