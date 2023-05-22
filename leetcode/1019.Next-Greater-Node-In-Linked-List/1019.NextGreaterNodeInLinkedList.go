package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	var nums []int
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	res, idxStack := make([]int, len(nums)), make([]int, 0, len(nums))
	for i, num := range nums {
		for len(idxStack) > 0 && num > nums[idxStack[len(idxStack)-1]] {
			idx := idxStack[len(idxStack)-1]
			res[idx] = num
			idxStack = idxStack[:len(idxStack)-1]
		}
		idxStack = append(idxStack, i)
	}
	return res
}

/*
[2,1,5]->[5,5,0]

[9,7,6,7,6,9]->[0,9,7,9,9,0]
*/
