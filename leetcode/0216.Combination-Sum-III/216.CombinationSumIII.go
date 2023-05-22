package leetcode

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	dfs(&ans, k, n, 1, []int{})
	return ans
}

func dfs(ans *[][]int, k, n, index int, path []int) {
	if n <= 0 {
		if n == 0 && k == len(path) {
			temp := make([]int, len(path))
			copy(temp, path)
			*ans = append(*ans, temp)
		}
		return
	}

	for i := index; i <= 9 && len(path)+1 <= k; i++ {
		if i > n {
			break
		}
		path = append(path, i)
		dfs(ans, k, n-i, i+1, path)
		path = path[:len(path)-1]
	}
}
