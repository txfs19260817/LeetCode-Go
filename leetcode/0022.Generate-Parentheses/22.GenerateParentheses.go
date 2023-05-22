package leetcode

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(int, int, string)
	dfs = func(lRemain, rRemain int, s string) {
		if lRemain == 0 && rRemain == 0 {
			ans = append(ans, s)
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, s+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, s+")")
		}
	}
	dfs(n, n, "")
	return ans
}
