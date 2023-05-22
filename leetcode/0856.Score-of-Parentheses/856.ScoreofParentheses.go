package leetcode

// https://leetcode.com/problems/score-of-parentheses/solution/314388
func scoreOfParentheses(s string) int {
	var ans, bal int
	for i, p := range s {
		if p == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' { // "(()(()))" -> "(()) + ((()))"
				ans += 1 << bal
			}
		}
	}
	return ans
}
