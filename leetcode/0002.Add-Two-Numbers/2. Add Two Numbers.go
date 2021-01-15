package _002_Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	var prevH *ListNode
	c := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			head.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			head.Val += l2.Val
			l2 = l2.Next
		}
		c = head.Val / 10
		head.Val %= 10
		head.Next = &ListNode{c, nil}
		prevH = head
		head = head.Next
	}
	if prevH != nil && head != nil && head.Val == 0 {
		prevH.Next = nil
	}
	return res
}

/*
[2,4,3]
[5,6,4]

[0]
[0]

[9,9,9,9,9,9,9]
[9,9,9,9]
*/
