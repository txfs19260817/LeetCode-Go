package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for p := head; p != nil; {
		newNode := &Node{p.Val, p.Next, p.Random}
		p.Next = newNode
		p = newNode.Next
	}
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	dummy := &Node{}
	q := dummy
	for p := head; p != nil; p = p.Next {
		q.Next = p.Next
		q = q.Next
		p.Next = q.Next
	}
	return dummy.Next
}
