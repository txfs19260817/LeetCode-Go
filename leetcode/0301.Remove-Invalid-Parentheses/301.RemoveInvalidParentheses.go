package _301_Remove_Invalid_Parentheses

func removeInvalidParentheses(s string) []string {
	var ans []string
	lRmv, rRmv := countMisplacedParentheses(s)
	var lCnt, rCnt int
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
			if str[i] == ')' {
				lCnt++
			} else if str[i] == ')' {
				rCnt++
			}
			if rCnt > lCnt { // pruning: invalid cases
				break
			}
		}
	}
	recursion(s, 0)
	return ans
}

func countMisplacedParentheses(s string) (l, r int) {
	for _, c := range s {
		switch c {
		case '(':
			l++
		case ')':
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return
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
