package _129_Sum_Root_to_Leaf_Numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	var path, nums []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil {
			var sum int
			for _, num := range path {
				sum = sum*10 + num
			}
			nums = append(nums, sum)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for _, num := range nums {
		ans += num
	}
	return ans
}
