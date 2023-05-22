package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			s = head
			for {
				if s == f {
					return s
				}
				s, f = s.Next, f.Next
			}
		}
	}
	return nil
}

/*
[3,2,0,-4]
1

[1,2]
0

[1]
-1
*/
