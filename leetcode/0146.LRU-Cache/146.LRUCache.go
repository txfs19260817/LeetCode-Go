package _146_LRU_Cache

type LRUCache struct {
	m          map[int]*listnode
	head, tail *listnode
	capacity   int
}

type listnode struct {
	key, value int
	prev, next *listnode
}

func (this *LRUCache) pushFront(node *listnode) {
	if node == nil {
		return
	}
	if this.head == nil {
		node.next, node.prev = nil, nil
		this.head = node
		this.tail = this.head
		return
	}
	node.next, node.prev = this.head, nil
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) remove(node *listnode) {
	if node == nil {
		return
	}
	if node == this.head {
		if this.head == this.tail {
			this.head, this.tail = nil, nil
			return
		}
		p := this.head.next
		this.head.next, p.prev = nil, nil
		this.head = p
		return
	}
	if node == this.tail {
		p := this.tail.prev
		this.head.prev, p.next = nil, nil
		this.tail = p
		return
	}
	after, before := node.next, node.prev
	after.prev = before
	before.next = after
}

func Constructor(capacity int) LRUCache {
	return LRUCache{m: make(map[int]*listnode), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.pushFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		node.value = value
		this.pushFront(node)
		return
	}
	node := &listnode{key: key, value: value}
	this.m[key] = node
	this.pushFront(node)
	if len(this.m) > this.capacity {
		k := this.tail.key
		this.remove(this.tail)
		delete(this.m, k)
	}
}
