package _290_Convert_Binary_Number_in_a_Linked_List_to_Integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for p := head; p != nil; p = p.Next {
		res = res*2 + p.Val
	}
	return res
}

func getDecimalValue1(head *ListNode) int {
	head = reverse(head)
	res, factor := 0, 1
	for p := head; p != nil; p = p.Next {
		res += p.Val * factor
		factor *= 2
	}
	return res
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	head.Next = nil
	for p != nil {
		temp := p.Next
		p.Next = head
		head = p
		p = temp
	}
	return head
}
