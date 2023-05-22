package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	m := map[int]*ListNode{}
	// 首次遍历建立 节点处链表和<->节点 哈希表
	// 若同一和出现多次会覆盖，即记录该sum出现的最后一次节点
	sum := 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		m[sum] = p
	}
	// 第二遍遍历 若当前节点处sum在下一处出现了则表明两结点之间所有节点和为0 直接删除区间所有节点
	sum = 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		p.Next = m[sum].Next
	}
	return dummy.Next
}
