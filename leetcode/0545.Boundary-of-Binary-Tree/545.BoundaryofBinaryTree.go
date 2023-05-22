package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	// root
	if !isLeaf(root) {
		ans = append(ans, root.Val)
	}
	//left Boundary
	for t := root.Left; t != nil && !isLeaf(t); {
		ans = append(ans, t.Val)
		if t.Left != nil {
			t = t.Left
		} else if t.Right != nil {
			t = t.Right
		}
	}
	//leaves
	addLeaves(&ans, root)
	//right Boundary
	var stack []int
	for t := root.Right; t != nil && !isLeaf(t); {
		stack = append(stack, t.Val)
		if t.Right != nil {
			t = t.Right
		} else if t.Left != nil {
			t = t.Left
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}

func isLeaf(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}

func addLeaves(ans *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*ans = append(*ans, root.Val)
		return
	}
	if root.Left != nil {
		addLeaves(ans, root.Left)
	}
	if root.Right != nil {
		addLeaves(ans, root.Right)
	}
}
