package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var x, y, prev *TreeNode
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > root.Val {
			y = root
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		prev = root
		root = root.Right
	}
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}
