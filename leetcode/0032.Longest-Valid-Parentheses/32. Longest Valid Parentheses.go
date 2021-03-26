package _032_Longest_Valid_Parentheses

func longestValidParentheses(s string) int {
	var ans, l, r int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r && l+r > ans {
			ans = l + r
		} else if r > l {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r && l+r > ans {
			ans = l + r
		} else if r < l {
			l, r = 0, 0
		}
	}
	return ans
}
