package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	folders := strings.Split(path, `/`)
	for _, folder := range folders {
		switch folder {
		case ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			if len(folder) > 0 {
				stack = append(stack, folder)
			}
		}
	}
	return `/` + strings.Join(stack, `/`)
}
