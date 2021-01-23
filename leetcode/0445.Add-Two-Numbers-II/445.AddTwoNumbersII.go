package _445_Add_Two_Numbers_II_

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{1, nil}
	len1, len2 := getLength(l1), getLength(l2)
	if len1 >= len2 {
		res.Next = add(l1, l2, len1-len2)
	} else {
		res.Next = add(l2, l1, len2-len1)
	}
	if res.Next != nil && res.Next.Val > 9 {
		res.Next.Val %= 10
		return res
	}
	return res.Next
}

func add(long, short *ListNode, offset int) *ListNode {
	if long == nil {
		return nil
	}
	var res *ListNode
	if offset == 0 {
		res = &ListNode{long.Val + short.Val, nil}
		res.Next = add(long.Next, short.Next, 0)
	} else {
		res = &ListNode{long.Val, nil}
		res.Next = add(long.Next, short, offset-1)
	}
	if res.Next != nil && res.Next.Val > 9 {
		res.Val++
		res.Next.Val %= 10
	}
	return res
}

func getLength(l *ListNode) int {
	count := 0
	for cur := l; cur != nil; cur = cur.Next {
		count++
	}
	return count
}
