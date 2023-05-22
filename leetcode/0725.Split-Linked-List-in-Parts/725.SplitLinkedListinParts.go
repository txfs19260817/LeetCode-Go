package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, 0, k)
	n := getLength(root)
	base, extra := n/k, n%k
	for i := 0; i < k; i++ {
		if root == nil {
			for ; i < k; i++ {
				res = append(res, nil)
			}
			return res
		}
		terms := base - 1
		if extra > 0 {
			terms++
			extra--
		}
		p := root
		for j := 0; p != nil && j < terms; j++ {
			p = p.Next
		}
		if p != nil {
			temp := p.Next
			p.Next = nil
			res = append(res, root)
			root = temp
		}
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

/*
[1,2,3,4]
5

[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
3
*/
