package _725_Split_Linked_List_in_Parts

import (
	"fmt"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	root := buildLinkedListFromInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	list := splitListToParts(root, 3)
	for _, node := range list {
		printLinkedList(node)
	}
}

func printLinkedList(head *ListNode) {
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

func buildLinkedListFromInts(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, num := range nums {
		node := &ListNode{num, nil}
		p.Next = node
		p = p.Next
	}
	return head.Next
}
