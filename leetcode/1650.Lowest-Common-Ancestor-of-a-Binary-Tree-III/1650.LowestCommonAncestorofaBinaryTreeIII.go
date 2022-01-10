package _650_Lowest_Common_Ancestor_of_a_Binary_Tree_III

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
