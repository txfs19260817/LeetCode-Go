package _988_Smallest_String_Starting_From_Leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var ans string
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[1:] }()
		if node.Left == nil && node.Right == nil {
			bs := make([]byte, 0, len(path))
			for i := len(path) - 1; i >= 0; i-- {
				bs = append(bs, byte(path[i]+'a'))
			}
			if s := string(bs); s < ans || len(ans) == 0 {
				ans = s
			}
			return
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, []int{})
	return ans
}
