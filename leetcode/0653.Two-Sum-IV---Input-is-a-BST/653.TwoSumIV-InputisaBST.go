package _653_Two_Sum_IV___Input_is_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var stack []*TreeNode
	var nums []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}
	for l, r := 0, len(nums)-1; l < r; {
		sum := nums[l] + nums[r]
		if sum == k {
			return true
		}
		if sum > k {
			r--
		} else {
			l++
		}
	}
	return false
}

func findTarget1(root *TreeNode, k int) bool {
	var stack []*TreeNode
	m := map[int]bool{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		root = root.Right
	}
	return false
}
