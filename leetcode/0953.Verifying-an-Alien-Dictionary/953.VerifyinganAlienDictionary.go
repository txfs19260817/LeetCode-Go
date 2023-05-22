package leetcode

func isAlienSorted(words []string, order string) bool {
	dict := map[byte]int{}
	for i, c := range order {
		dict[byte(c)] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			// if both words share the same prefix, the latter must longer than or equal to the former
			if j >= len(words[i+1]) {
				return false
			}
			// the corresponding position must in alien dict order
			if words[i][j] != words[i+1][j] {
				if dict[words[i][j]] > dict[words[i+1][j]] {
					return false
				}
				break // the remain part is not to be considered
			}
		}
	}
	return true
}
