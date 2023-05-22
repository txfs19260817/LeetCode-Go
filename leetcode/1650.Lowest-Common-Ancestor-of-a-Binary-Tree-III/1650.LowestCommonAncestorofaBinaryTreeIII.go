package leetcode

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	a, b := p, q
	for a != b {
		if a == nil {
			a = p
		} else {
			a = a.Parent
		}
		if b == nil {
			b = q
		} else {
			b = b.Parent
		}
	}
	return a
}
