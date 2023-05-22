package leetcode

func countArrangement(n int) int {
	var ans int
	visited := make([]bool, n+1)
	dfs(n, 1, &ans, visited)
	return ans
}

func dfs(n, pos int, count *int, visited []bool) {
	if n < pos {
		*count++
		return
	}
	for i := 1; i <= n; i++ {
		if !visited[i] && (i%pos == 0 || pos%i == 0) {
			visited[i] = true
			dfs(n, pos+1, count, visited)
			visited[i] = false
		}
	}
}
