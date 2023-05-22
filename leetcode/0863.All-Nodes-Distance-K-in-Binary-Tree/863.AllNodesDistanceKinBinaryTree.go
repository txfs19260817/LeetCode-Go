package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int
	parents := map[int]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)
	var findAns func(node, from *TreeNode, depth int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return ans
}

func distanceK2(root *TreeNode, target *TreeNode, K int) []int {
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
