package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy, ans := head, head
	head = head.Next
	dummy.Next = nil

	for head != nil {
		var sum int
		for {
			val := head.Val
			sum += val
			head = head.Next
			if val == 0 {
				break
			}
		}
		dummy.Next = &ListNode{Val: sum}
		dummy = dummy.Next
	}
	return ans.Next
}
