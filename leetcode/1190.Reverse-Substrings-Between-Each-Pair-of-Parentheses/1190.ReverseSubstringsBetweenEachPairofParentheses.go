package _190_Reverse_Substrings_Between_Each_Pair_of_Parentheses

func reverseParentheses(s string) string {
	pair, stack, ans := make([]int, len(s)), make([]int, 0, len(s)), make([]byte, 0, len(s))
	for i, c := range s {
		switch c {
		case '(':
			stack = append(stack, i)
		case ')':
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}
	for i, dir := 0, 1; i < len(s); i += dir {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			dir *= -1
			continue
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}
