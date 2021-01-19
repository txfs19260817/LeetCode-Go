package _141_Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}

/*
[3,2,0,-4]
1

[1,2]
0

[1]
-1
*/
