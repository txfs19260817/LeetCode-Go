package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{} // create new head
	curr, prev := newHead, head
	for prev != nil {
		for curr.Next != nil && curr.Next.Val < prev.Val {
			curr = curr.Next // find the insertion position in the new sequence
		}
		temp := prev.Next     // save the next to-be-processed element pointer
		prev.Next = curr.Next // insert 1
		curr.Next = prev      // insert 2
		prev = temp           // place `prev` to the next old element
		curr = newHead        // return to the new head
	}
	return newHead.Next // remember to remove the new head from the return value
}
