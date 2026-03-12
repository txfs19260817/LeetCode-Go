package leetcode

// findNeedleInChunkedStrings treats chunks as one continuous string and returns
// the start position of the first match as [chunkIndex, offset].
func findNeedleInChunkedStrings(chunks []string, needle string) [2]int {
	if len(needle) == 0 {
		return [2]int{0, 0}
	}

	nextPos := func(i, j int) (int, int) {
		j++
		for i < len(chunks) && j >= len(chunks[i]) {
			i++
			j = 0
		}
		return i, j
	}
	for i := 0; i < len(chunks); i++ {
		for j := 0; j < len(chunks[i]); j++ {
			k := 0
			x, y := i, j
			for k < len(needle) && x < len(chunks) {
				if needle[k] != chunks[x][y] {
					break
				}
				k++
				if k == len(needle) {
					return [2]int{i, j}
				}
				x, y = nextPos(x, y)
			}
		}
	}
	return [2]int{-1, -1}
}
