package waymo

type avlNode struct {
	key    int
	height int
	left   *avlNode
	right  *avlNode
}

type AVLTree struct {
	root *avlNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(key int) {
	t.root = insertNode(t.root, key)
}

func (t *AVLTree) Delete(key int) {
	t.root = deleteNode(t.root, key)
}

func (t *AVLTree) Contains(key int) bool {
	cur := t.root
	for cur != nil {
		if key < cur.key {
			cur = cur.left
			continue
		}
		if key > cur.key {
			cur = cur.right
			continue
		}
		return true
	}
	return false
}

func (t *AVLTree) InOrder() []int {
	out := make([]int, 0)
	var walk func(*avlNode)
	walk = func(node *avlNode) {
		if node == nil {
			return
		}
		walk(node.left)
		out = append(out, node.key)
		walk(node.right)
	}
	walk(t.root)
	return out
}

func insertNode(node *avlNode, key int) *avlNode {
	if node == nil {
		return &avlNode{key: key, height: 1}
	}
	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}
	return rebalance(node)
}

func deleteNode(node *avlNode, key int) *avlNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = deleteNode(node.left, key)
		return rebalance(node)
	}
	if key > node.key {
		node.right = deleteNode(node.right, key)
		return rebalance(node)
	}

	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}

	successor := minNode(node.right)
	node.key = successor.key
	node.right = deleteNode(node.right, successor.key)
	return rebalance(node)
}

func minNode(node *avlNode) *avlNode {
	cur := node
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func rebalance(node *avlNode) *avlNode {
	updateHeight(node)
	balance := height(node.left) - height(node.right)

	if balance > 1 {
		if height(node.left.left) < height(node.left.right) {
			node.left = rotateLeft(node.left)
		}
		return rotateRight(node)
	}
	if balance < -1 {
		if height(node.right.right) < height(node.right.left) {
			node.right = rotateRight(node.right)
		}
		return rotateLeft(node)
	}
	return node
}

func rotateLeft(x *avlNode) *avlNode {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	updateHeight(x)
	updateHeight(y)
	return y
}

func rotateRight(y *avlNode) *avlNode {
	x := y.left
	t2 := x.right

	x.right = y
	y.left = t2

	updateHeight(y)
	updateHeight(x)
	return x
}

func updateHeight(node *avlNode) {
	node.height = max(height(node.left), height(node.right)) + 1
}

func height(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
