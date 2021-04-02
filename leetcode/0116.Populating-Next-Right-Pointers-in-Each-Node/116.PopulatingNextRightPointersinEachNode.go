package _116_Populating_Next_Right_Pointers_in_Each_Node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for i := 0; len(q) > 0; i++ {
		p := make([]*Node, 0, 1<<(i+1))
		for j := 0; j < len(q); j++ {
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		for k := 0; k < len(p)-1; k++ {
			p[k].Next = p[k+1]
		}
		q = p
	}
	return root
}
