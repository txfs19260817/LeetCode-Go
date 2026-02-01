package affirm

func PrefixToPostfix(prefix string) string {
	if len(prefix) <= 1 {
		return prefix
	}

	stack := make([]string, 0, len(prefix))
	for i := len(prefix) - 1; i >= 0; i-- {
		ch := prefix[i]
		if isOperator(ch) {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, left+right+string(ch))
		} else {
			stack = append(stack, string(ch))
		}
	}
	return stack[len(stack)-1]
}

func isOperator(ch byte) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}
