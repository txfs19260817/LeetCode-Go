package leetcode

func imageSmoother(img [][]int) [][]int {
	ans := make([][]int, len(img))
	for i := range img {
		ans[i] = make([]int, len(img[i]))
	}
	for i := range img {
		for j := range img[i] {
			var cnt int
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if r >= 0 && r < len(img) && c >= 0 && c < len(img[0]) {
						ans[i][j] += img[r][c]
						cnt++
					}
				}
			}
			ans[i][j] /= cnt
		}
	}
	return ans
}
