package k

func one(mat [][]byte) (start, end []int, height, width int) {
	for i, row := range mat {
		for j, v := range row {
			if v == '0' {
				start = []int{i, j}
				goto end
			}
		}
	}
end:
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] == '0' {
				end = []int{i, j}
				goto ret
			}
		}
	}
ret:
	height, width = end[0]-start[0]+1, end[1]-start[1]+1
	return
}

func multiple(mat [][]byte) (starts, ends [][]int) {
	search := func(i, j int) []int {
		c := j
		for ; c < len(mat[i]); c++ {
			if mat[i][c] != '0' {
				c--
				break
			}
		}
		if c == len(mat[i]) {
			c = len(mat[i]) - 1
		}
		r := i
		for ; r < len(mat); r++ {
			if mat[r][c] != '0' {
				r--
				break
			}
		}
		if r == len(mat) {
			r = len(mat) - 1
		}
		// mark the rect as visited
		for p := i; p <= r; p++ {
			for q := j; q <= c; q++ {
				mat[p][q] = 'x'
			}
		}
		return []int{r, c}
	}
	for i, row := range mat {
		for j, v := range row {
			if v == '0' {
				starts = append(starts, []int{i, j})
				ends = append(ends, search(i, j))
			}
		}
	}
	return
}
