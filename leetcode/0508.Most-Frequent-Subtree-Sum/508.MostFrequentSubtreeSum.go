package _508_Most_Frequent_Subtree_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	freq := map[int]int{}
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			freq[node.Val]++
			return node.Val
		}
		sum := node.Val + helper(node.Left) + helper(node.Right)
		freq[sum]++
		return sum
	}
	helper(root)
	var maxFreq int
	for _, v := range freq {
		if maxFreq < v {
			maxFreq = v
		}
	}
	for k, v := range freq {
		if maxFreq == v {
			ans = append(ans, k)
		}
	}
	return ans
}
