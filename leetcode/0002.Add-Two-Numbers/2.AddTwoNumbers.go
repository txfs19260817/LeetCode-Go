package _002_Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur, pre := head, head
	var c int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			cur.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}
		c = cur.Val / 10
		cur.Val %= 10
		cur.Next = &ListNode{c, nil}
		pre = cur
		cur = cur.Next
	}
	if cur.Val == 0 {
		pre.Next = nil
	}
	return head
}
