package leetcode

type Node struct {
	Val   int
	Left  *Node // prev
	Right *Node // next
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	dummy := &Node{}
	p := dummy
	var stack []*Node
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		p.Right = root
		root.Left = p
		p = p.Right

		root = root.Right
	}

	ans := dummy.Right // the "head"
	p.Right = ans
	ans.Left = p
	return ans
}
