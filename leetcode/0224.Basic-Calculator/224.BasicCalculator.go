package _224_Basic_Calculator

func calculate(s string) int {
	ans, sign := 0, 1
	var stack []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, ans)  // the 2nd is cached ans
			stack = append(stack, sign) // top is sign
			ans, sign = 0, 1
		case ')':
			ans = ans*stack[len(stack)-1] + stack[len(stack)-2] // ans*sign+cached
			stack = stack[:len(stack)-2]
		default: // number byte
			num := int(s[i] - '0')
			for ; i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9'; i++ {
				num = num*10 + int(s[i+1]-'0')
			}
			ans += num * sign
		}
	}
	return ans
}
