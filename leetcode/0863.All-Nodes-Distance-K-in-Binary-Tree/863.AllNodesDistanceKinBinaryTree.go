package _863_All_Nodes_Distance_K_in_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	var ans []int
	dfs(root, target, K, &ans)
	return ans
}

func dfs(root *TreeNode, target *TreeNode, K int, ans *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		findChildren(root, K, ans)
		return K - 1
	}
	leftDistance := dfs(root.Left, target, K, ans)
	if leftDistance == 0 {
		findChildren(root, 0, ans)
	}
	if leftDistance > 0 {
		findChildren(root.Right, leftDistance-1, ans)
		return leftDistance - 1
	}
	rightDistance := dfs(root.Right, target, K, ans)
	if rightDistance == 0 {
		findChildren(root, 0, ans)
	}
	if rightDistance > 0 {
		findChildren(root.Left, rightDistance-1, ans)
		return rightDistance - 1
	}
	return -1
}

func findChildren(root *TreeNode, K int, ans *[]int) {
	if root == nil {
		return
	}
	if K == 0 {
		*ans = append(*ans, root.Val)
	} else {
		findChildren(root.Left, K-1, ans)
		findChildren(root.Right, K-1, ans)
	}
}
