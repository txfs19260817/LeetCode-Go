package leetcode

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for t := 0; t < k; t++ {
			if mat1[i][t] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if mat2[t][j] != 0 {
					result[i][j] += mat1[i][t] * mat2[t][j]
				}
			}
		}
	}
	return result
}
