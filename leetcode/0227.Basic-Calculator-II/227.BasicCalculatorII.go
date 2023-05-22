package leetcode

func calculate(s string) int {
	ans, curNum, preNum, sign := 0, 0, 0, '+'
	for i, c := range s {
		if c >= '0' && c <= '9' {
			curNum = curNum*10 + int(c-'0')
		}
		if c == '+' || c == '-' || c == '*' || c == '/' || i == len(s)-1 {
			switch sign {
			case '+':
				ans += preNum
				preNum = curNum
			case '-':
				ans += preNum
				preNum = -1 * curNum
			case '*':
				preNum *= curNum
			case '/':
				preNum /= curNum
			}
			curNum, sign = 0, c
		}
	}
	return ans + preNum
}
