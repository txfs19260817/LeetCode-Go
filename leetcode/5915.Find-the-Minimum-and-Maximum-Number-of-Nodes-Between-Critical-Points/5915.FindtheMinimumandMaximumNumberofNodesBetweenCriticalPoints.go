package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	firstIdx, lastIdx, minDistance := 0, 0, 1000000
	for pre, cur, idx := head, head.Next, 1; cur.Next != nil; cur = cur.Next {
		if cur.Val > pre.Val && cur.Val > cur.Next.Val || cur.Val < pre.Val && cur.Val < cur.Next.Val {
			if firstIdx == 0 {
				firstIdx = idx
			} else {
				if lastIdx == 0 {
					minDistance = idx - firstIdx
				} else if v := idx - lastIdx; minDistance > v {
					minDistance = v
				}
				lastIdx = idx
			}
		}
		pre = cur
		idx++
	}
	if firstIdx == 0 || lastIdx == 0 {
		return []int{-1, -1}
	}
	return []int{minDistance, lastIdx - firstIdx}
}
