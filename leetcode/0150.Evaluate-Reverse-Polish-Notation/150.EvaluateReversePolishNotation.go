package _150_Evaluate_Reverse_Polish_Notation

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			default:
				stack = append(stack, a/b)
			}
		}
	}
	return stack[0]
}
