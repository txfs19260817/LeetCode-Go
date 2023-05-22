package leetcode

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode = &Node{x, nil}
		aNode.Next = aNode
		return aNode
	}
	if aNode.Next == aNode {
		temp := &Node{x, aNode}
		aNode.Next = temp
		return aNode
	}
	p, q := aNode, aNode.Next
	maxNode := aNode
	for {
		p, q = p.Next, q.Next
		if p.Val <= x && q.Val >= x {
			temp := &Node{x, q}
			p.Next = temp
			return aNode
		}
		if p.Val >= maxNode.Val {
			maxNode = p
		}
		if p == aNode {
			break
		}
	}
	temp := &Node{x, maxNode.Next}
	maxNode.Next = temp
	return aNode
}

/*
[3,4,1]
2

[1]
0

[]
1

[3,3,3]
0

[1,3,5]
0
*/
