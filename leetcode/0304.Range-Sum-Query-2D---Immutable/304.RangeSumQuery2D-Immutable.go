package leetcode

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefix := make([][]int, len(matrix)+1)
	for i := range prefix {
		prefix[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(prefix); i++ {
		for j := 1; j < len(prefix[0]); j++ {
			prefix[i][j] = matrix[i-1][j-1] + prefix[i][j-1] + prefix[i-1][j] - prefix[i-1][j-1]
		}
	}
	return NumMatrix{prefix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prefix[row2+1][col2+1] - this.prefix[row1][col2+1] - this.prefix[row2+1][col1] + this.prefix[row1][col1]
}
