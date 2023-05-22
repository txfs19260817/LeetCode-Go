package leetcode

func checkValidString(s string) bool {
	return helper(s, false) && helper(s, true)
}

func helper(s string, reverse bool) bool {
	var n, star, i int
	var l, r byte = '(', ')'
	if reverse {
		l, r = r, l
		i = len(s) - 1
	}
	traverse := func(c byte) bool {
		switch c {
		case '*':
			star++
		case l:
			n++
		case r:
			if n == 0 {
				if star == 0 {
					return false
				}
				star--
			} else {
				n--
			}
		}
		return true
	}
	for i >= 0 && i < len(s) {
		if earlyRet := traverse(s[i]); !earlyRet {
			return false
		}
		if reverse {
			i--
		} else {
			i++
		}
	}

	return n == 0 || n <= star
}

func checkValidString2(s string) bool {
	var maxLeft, minLeft int
	for _, c := range s {
		switch c {
		case '(':
			maxLeft++
			minLeft++
		case ')':
			maxLeft--
			if maxLeft < 0 {
				return false
			}
			minLeft = max(minLeft-1, 0)
		case '*':
			maxLeft++
			minLeft = max(minLeft-1, 0)
		}
	}
	return minLeft == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
