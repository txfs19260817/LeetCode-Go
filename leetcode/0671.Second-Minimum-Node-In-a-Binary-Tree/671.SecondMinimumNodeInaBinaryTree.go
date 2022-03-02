package _671_Second_Minimum_Node_In_a_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	mini, ans, q := root.Val, max(root.Left.Val, root.Right.Val), []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Val > ans && ans > mini {
				continue
			}
			if node.Val > mini && (node.Val < ans || ans == mini) {
				ans = node.Val
			}
			if node.Left != nil {
				p = append(p, node.Left)
				p = append(p, node.Right)
			}
		}
		q = p
	}
	if mini == ans {
		return -1
	}
	return ans
}

func findSecondMinimumValue2(root *TreeNode) int {
	ans := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
