package _498_Diagonal_Traverse

func findDiagonalOrder(mat [][]int) []int {
	ans := make([]int, 0, len(mat)*len(mat[0]))
	for sum := 0; sum < len(mat)+len(mat[0])-1; sum++ {
		if sum%2 == 1 {
			for i := 0; i <= sum; i++ {
				if j := sum - i; i < len(mat) && j < len(mat[0]) {
					ans = append(ans, mat[i][j])
				}
			}
		} else {
			for j := 0; j <= sum; j++ {
				if i := sum - j; i < len(mat) && j < len(mat[0]) {
					ans = append(ans, mat[i][j])
				}
			}
		}
	}
	return ans
}
