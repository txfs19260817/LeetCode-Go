package _669_Merge_In_Between_Linked_Lists_

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p := list1
	for i := 0; i < a-1; i++ {
		p = p.Next
	}
	befA := p
	for i := 0; i < b-a+1; i++ {
		p = p.Next
	}
	aftB := p.Next
	p = list2
	for p.Next != nil {
		p = p.Next
	}
	befA.Next, p.Next = list2, aftB
	return list1
}
