package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := map[int]bool{}
	for _, k := range G {
		m[k] = true
	}
	res := 0
	for ; head != nil; head = head.Next {
		if m[head.Val] {
			for head.Next != nil && m[head.Next.Val] {
				head = head.Next
			}
			res++
		}
	}
	return res
}

/*
[0,1,2,3]
[0,1,3]

[0,1,2,3,4]
[0,3,1,4]
*/
