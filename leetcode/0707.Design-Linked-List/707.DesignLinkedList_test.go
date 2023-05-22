package leetcode

import (
	"fmt"
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	l := Constructor()
	l.AddAtHead(7)
	printLinkedList(&l, "AddAtHead 7")
	l.AddAtHead(2)
	printLinkedList(&l, "AddAtHead 2")
	l.AddAtHead(1)
	printLinkedList(&l, "AddAtHead 1")
	l.AddAtIndex(3, 0)
	printLinkedList(&l, "AddAtIndex 3,0")
	l.DeleteAtIndex(2)
	printLinkedList(&l, "DeleteAtIndex 2")
	l.AddAtHead(6)
	printLinkedList(&l, "AddAtHead 6")
	l.AddAtTail(4)
	printLinkedList(&l, "AddAtTail 4")
	fmt.Println("Get 4: ", l.Get(4))
	l.AddAtHead(4)
	printLinkedList(&l, "AddAtHead 4")
	l.AddAtIndex(5, 0)
	printLinkedList(&l, "AddAtIndex 5,0")
	l.AddAtHead(6)
	printLinkedList(&l, "AddAtHead 6")
}

func printLinkedList(head *MyLinkedList, info ...string) {
	if len(info) != 0 {
		fmt.Print(info[0] + ": ")
	}
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
