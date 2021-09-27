package _430_Flatten_a_Multilevel_Doubly_Linked_List

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	for p := root; p != nil; p = p.Next {
		if p.Child == nil {
			continue
		}
		childEnd := p.Child
		for childEnd.Next != nil {
			childEnd = childEnd.Next
		}
		childEnd.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = childEnd
		}
		p.Next = p.Child
		p.Child.Prev = p
		p.Child = nil
	}
	return root
}
