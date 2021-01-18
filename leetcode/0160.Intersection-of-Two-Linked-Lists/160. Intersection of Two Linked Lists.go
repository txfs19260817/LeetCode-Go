package _160_Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	for p := headA; p != nil; p = p.Next {
		for q := headB; q != nil; q = q.Next {
			if p == q {
				return p
			}
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}

	p := headA
	for p.Next != nil {
		p = p.Next
	}
	p.Next = headA

	s, f := headB, headB
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			s = headB
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
8
[4,1,8,4,5]
[5,6,1,8,4,5]
2
3

2
[1,9,1,2,4]
[3,2,4]
3
1

0
[2,6,4]
[1,5]
3
2

3
[3]
[2,3]
0
1
*/
