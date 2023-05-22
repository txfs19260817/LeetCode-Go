package leetcode

func addOperators(num string, target int) []string {
	var ans []string
	var recursion func(expr []byte, start, res, mul int)
	recursion = func(expr []byte, start, res, mul int) {
		if len(num) == start {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		// create a placeholder for filling the operator
		opIdx := len(expr)
		if start > 0 {
			expr = append(expr, 0)
		}
		for i, val := start, 0; i < len(num) && (i == start || num[start] != '0'); i++ { // no leading 0
			val = val*10 + int(num[i]-'0')
			expr = append(expr, num[i])
			if start == 0 {
				recursion(expr, i+1, val, val)
				continue
			}
			expr[opIdx] = '+'
			recursion(expr, i+1, res+val, val)
			expr[opIdx] = '-'
			recursion(expr, i+1, res-val, -val)
			expr[opIdx] = '*'
			recursion(expr, i+1, res-mul+mul*val, mul*val)
		}
	}
	recursion(make([]byte, 0, 2*len(num)-1), 0, 0, 0)
	return ans
}
