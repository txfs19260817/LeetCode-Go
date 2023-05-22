package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return builder(nums, 0, len(nums)-1)
}

func builder(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = builder(nums, l, mid-1)
	root.Right = builder(nums, mid+1, r)
	return root
}
