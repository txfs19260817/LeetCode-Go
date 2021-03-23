# 150. Evaluate Reverse Polish Notation

## [LeetCode **150. Evaluate Reverse Polish Notation**](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

### Description

Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, \*, and /. Each operand may be an integer or another expression.

Note that division between two integers should truncate toward zero.

It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.

Example 1:

```text
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
```

Example 2:

```go
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
```

Example 3:

```go
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

Constraints:

* `1 <= tokens.length <= 10^4`
* `tokens[i]`is either an operator: `"+"`, `"-"`, `"*"`, or `"/"`, or an integer in the range `[-200, 200]`.

### Tags

Stack

### Solution

使用栈解决。遍历输入数组，如果是数字则直接将其入栈，否则就是操作符，记作`o`。依次弹出栈顶两个元素，记作`b`和`a`，然后将计算结果`a o b`入栈。最后返回栈中唯一所剩元素。

A stack is applied here. Traverse every token in the input array. If the token is a number, then add it as an integer to the stack. Otherwise it should be an operator and let `o`  denote it. Pop 2 elements from the stack successively and denote them `b` and `a` respectively. Then push the result of `a o b` into the stack. Finally, return the only remaining element in this stack.

### Code

```go
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
```

## Reference

1. [逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/solution/ni-bo-lan-biao-da-shi-qiu-zhi-by-leetcod-wue9/)

