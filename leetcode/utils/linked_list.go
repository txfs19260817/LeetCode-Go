package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedListFromInts(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, num := range nums {
		node := &ListNode{num, nil}
		p.Next = node
		p = p.Next
	}
	return head.Next
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		} else {
			fmt.Println()
		}
		head = head.Next
	}
}

func AreLinkedListsEqual(l1, l2 *ListNode) bool {
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}
