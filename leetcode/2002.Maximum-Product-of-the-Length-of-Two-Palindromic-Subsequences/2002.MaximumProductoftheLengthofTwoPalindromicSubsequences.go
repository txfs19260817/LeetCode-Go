package leetcode

func maxProduct(s string) int {
	var ans int
	palindromic := func(s []byte) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[r] != s[l] {
				return false
			}
		}
		return true
	}
	var dfs func(s1, s2 []byte, index int)
	dfs = func(s1, s2 []byte, index int) {
		if palindromic(s1) && palindromic(s2) {
			if v := len(s1) * len(s2); ans < v {
				ans = v
			}
		}
		if index == len(s) {
			return
		}
		dfs(append(s1, s[index]), s2, index+1)
		dfs(s1, append(s2, s[index]), index+1)
		dfs(s1, s2, index+1)
	}
	dfs([]byte{}, []byte{}, 0)
	return ans
}
