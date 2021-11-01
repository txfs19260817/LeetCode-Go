package _060_Check_if_an_Original_String_Exists_Given_Two_Encoded_Strings

func possiblyEquals(s1 string, s2 string) bool {
	visited, bias := make([][][2000]bool, len(s1)+1), 1000 // diff range (-1000, 1000)
	for i := range visited {
		visited[i] = make([][2000]bool, len(s2)+1)
	}
	isNum := func(b byte) bool { return b >= '1' && b <= '9' } // no '0' in this problem
	var dfs func(i, j, d int) bool
	dfs = func(i, j, d int) bool {
		if i == len(s1) && j == len(s2) {
			return d == 0 // matched
		}
		if visited[i][j][d+bias] {
			return false // visited
		}
		visited[i][j][d+bias] = true
		if d == 0 && i < len(s1) && j < len(s2) && s1[i] == s2[j] && dfs(i+1, j+1, d) { // len(origin_s1) == len(origin_s2), extend both
			return true
		}
		if d <= 0 && i < len(s1) { // len(origin_s1) <= len(origin_s2), extend s1
			if isNum(s1[i]) { // s1[i] is numeric
				for k, l := i, 0; k < len(s1) && isNum(s1[k]); k++ {
					l = l*10 + int(s1[k]&15)
					if dfs(k+1, j, d+l) {
						return true
					}
				}
			} else if d < 0 && dfs(i+1, j, d+1) { // s1[i] is alphabet, note that it should be len(origin_s1) < len(origin_s2) so that extend s1 by 1 is meaningful
				return true
			}
		}
		if d >= 0 && j < len(s2) { // len(origin_s1) >= len(origin_s2), extend s2
			if isNum(s2[j]) { // s2[j] is numeric
				for k, l := j, 0; k < len(s2) && isNum(s2[k]); k++ {
					l = l*10 + int(s2[k]&15)
					if dfs(i, k+1, d-l) {
						return true
					}
				}
			} else if d > 0 && dfs(i, j+1, d-1) { // s2[j] is alphabet, it should be len(origin_s1) > len(origin_s2) so that extend s2 by 1 is meaningful
				return true
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
