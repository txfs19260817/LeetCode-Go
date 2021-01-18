package _000_playground

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := middleOfLL(head)
	temp:=mid.Next
	mid.Next=nil
	mid=temp
	return mergeSort(sortList(head), sortList(mid))
}

func mergeSort(left *ListNode, right *ListNode) *ListNode {
	head := &ListNode{}
	res:=head
	for left!=nil && right!=nil {
		if left.Val<right.Val{
			head.Next=left
			left=left.Next
		} else {
			head.Next=right
			right=right.Next
		}
		head=head.Next
		head.Next=nil
	}
	if left != nil {
		head.Next=left
	}
	if right!=nil {
		head.Next=right
	}
	return res.Next
}

func middleOfLL(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s,f:=head,head.Next
	for f.Next!=nil&&f.Next.Next!=nil {
		s,f = s.Next, f.Next.Next
	}
	return s
}