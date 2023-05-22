package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	var j int
	stack := make([]int, 0, len(pushed))
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0 && j == len(popped)
}
