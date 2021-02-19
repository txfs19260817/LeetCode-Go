package _131_Palindrome_Partitioning

func partition(s string) [][]string {
	var ans [][]string
	dfs(&ans, s, []string{})
	return ans
}

func dfs(ans *[][]string, s string, path []string) {
	if len(s) == 0 {
		temp := make([]string, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			path = append(path, s[:i])
			dfs(ans, s[i:], path)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
