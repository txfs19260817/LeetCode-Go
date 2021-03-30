package _445_Add_Two_Numbers_II_

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	len1, len2 := getLength(l1), getLength(l2)
	if len1 > len2 {
		head = recursion(l1, l2, len1-len2)
	} else {
		head = recursion(l2, l1, len2-len1)
	}
	if head != nil && head.Val > 9 {
		temp := &ListNode{head.Val / 10, head}
		head.Val %= 10
		head = temp
	}
	return head
}

func recursion(long, short *ListNode, delta int) *ListNode {
	if long == nil {
		return nil
	}
	curNode := &ListNode{}
	if delta == 0 {
		curNode.Val = long.Val + short.Val
		curNode.Next = recursion(long.Next, short.Next, 0)
	} else {
		curNode.Val = long.Val
		curNode.Next = recursion(long.Next, short, delta-1)
	}
	if curNode.Next != nil && curNode.Next.Val > 9 {
		curNode.Val += curNode.Next.Val / 10
		curNode.Next.Val %= 10
	}
	return curNode
}

func getLength(head *ListNode) int {
	var l int
	for p := head; p != nil; p = p.Next {
		l++
	}
	return l
}
