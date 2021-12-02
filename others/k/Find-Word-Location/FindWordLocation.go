package k

func findWordLocation(grid [][]byte, word string) [][]int {
	for i := range grid {
		for j := range grid[i] {
			if ans := search(grid, word, 0, i, j); ans != nil {
				return ans
			}
		}
	}
	return nil
}

func search(grid [][]byte, word string, index, i, j int) [][]int {
	if i >= len(grid) || j >= len(grid[0]) || word[index] != grid[i][j] {
		return nil
	}
	if index == len(word)-1 {
		return [][]int{{i, j}}
	}
	next := search(grid, word, index+1, i, j+1) // right
	if next == nil {
		next = search(grid, word, index+1, i+1, j) // down
		if next == nil {
			return nil
		}
	}
	return append([][]int{{i, j}}, next...)
}
