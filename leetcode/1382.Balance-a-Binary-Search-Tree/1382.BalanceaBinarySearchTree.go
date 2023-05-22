package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	return slice2bst(inorder(root))
}

func inorder(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}

func slice2bst(nums []int) *TreeNode {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nums[0]}
	default:
		return &TreeNode{
			Val:   nums[len(nums)/2],
			Left:  slice2bst(nums[:len(nums)/2]),
			Right: slice2bst(nums[len(nums)/2+1:]),
		}

	}
}
