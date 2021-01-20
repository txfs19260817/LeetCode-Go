package _237_Delete_Node_in_a_Linked_List_

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

/*
[4,5,1,9]
5

[-3,5,-99]
-3
*/
