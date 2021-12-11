package _301_Remove_Invalid_Parentheses

func removeInvalidParentheses(s string) []string {
	var ans []string
	var lRmv, rRmv int // count misplaced parentheses that to be removed
	for _, c := range s {
		switch c {
		case '(':
			lRmv++
		case ')':
			if lRmv > 0 {
				lRmv--
			} else {
				rRmv++
			}
		}
	}
	var recursion func(str string, idx int)
	recursion = func(str string, idx int) {
		if lRmv == 0 && rRmv == 0 {
			if isValid(str) {
				ans = append(ans, str)
			}
			return
		}
		for i := idx; i < len(str); i++ {
			if i != idx && str[i] == str[i-1] { // pruning: dedupe
				continue
			}
			if lRmv+rRmv > len(str)-idx { // pruning: no enough chars to remove
				break
			}
			if lRmv > 0 && str[i] == '(' {
				lRmv--
				recursion(str[:i]+str[i+1:], i)
				lRmv++
			} else if rRmv > 0 && str[i] == ')' {
				rRmv--
				recursion(str[:i]+str[i+1:], i)
				rRmv++
			}
		}
	}
	recursion(s, 0)
	return ans
}

func isValid(s string) bool {
	var cnt int
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
