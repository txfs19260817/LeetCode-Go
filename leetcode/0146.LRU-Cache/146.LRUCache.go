package _146_LRU_Cache

type LRUCache struct {
	cap        int
	m          map[int]*listNode
	head, tail *listNode
}

type listNode struct {
	key, value int // store key to delete tail node from m easily
	prev, next *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, m: make(map[int]*listNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		if node != this.head {
			// move to head
			this.removeNode(node)
			this.addNode(node)
		}
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.removeNode(node)
		node.value = value
		this.addNode(node)
		return
	}
	newNode := &listNode{key: key, value: value}
	this.addNode(newNode)
	this.m[key] = newNode
	if len(this.m) > this.cap {
		delete(this.m, this.tail.key)
		this.removeNode(this.tail)
	}
}

func (this *LRUCache) addNode(node *listNode) {
	node.next, node.prev = this.head, nil
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = this.head
		this.tail.next = nil
	}
}

func (this *LRUCache) removeNode(node *listNode) {
	if node == nil {
		return
	}
	if node == this.head {
		this.head = node.next
		node.next = nil
		if this.head != nil { // in case there is only 1 node in list
			this.head.prev = nil
		}
		return
	}
	if node == this.tail {
		this.tail = node.prev
		node.prev = nil
		this.tail.next = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}
