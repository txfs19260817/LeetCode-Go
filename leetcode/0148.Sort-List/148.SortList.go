package _148_Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := middleNode(head)
	temp := mid.Next
	mid.Next = nil
	mid = temp
	return merge(sortList(head), sortList(mid))
}

func merge(left, right *ListNode) *ListNode {
	res := &ListNode{}
	head := res

	for left != nil && right != nil {
		if left.Val < right.Val {
			res.Next = left
			left = left.Next
		} else {
			res.Next = right
			right = right.Next
		}
		res = res.Next
		res.Next = nil
	}
	if left != nil {
		res.Next = left
	}
	if right != nil {
		res.Next = right
	}
	return head.Next
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	if fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
