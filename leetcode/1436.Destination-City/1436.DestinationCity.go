package leetcode

func destCity(paths [][]string) string {
	ans, m := "", map[string]bool{}
	for _, path := range paths {
		m[path[0]] = true
	}
	for _, path := range paths {
		if !m[path[1]] {
			ans = path[1]
			break
		}
	}
	return ans
}
