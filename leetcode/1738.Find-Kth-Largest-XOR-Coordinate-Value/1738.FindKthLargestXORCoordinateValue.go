package _738_Find_Kth_Largest_XOR_Coordinate_Value

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	results, pre := make([]int, 0, len(matrix)*len(matrix[0])), make([][]int, len(matrix)+1)
	pre[0] = make([]int, len(matrix[0])+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, len(matrix[0])+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return results[k-1]
}
