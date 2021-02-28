package _255_Verify_Preorder_Sequence_in_Binary_Search_Tree

func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}
	l, r := 1, 0
	for i, n := range preorder {
		if n > preorder[0] {
			r = i
			break
		}
	}
	if r == 0 {
		return verifyPreorder(preorder[1:])
	}
	for i := r; i < len(preorder); i++ {
		if preorder[i] < preorder[0] {
			return false
		}
	}
	return verifyPreorder(preorder[l:r]) && verifyPreorder(preorder[r:])
}
