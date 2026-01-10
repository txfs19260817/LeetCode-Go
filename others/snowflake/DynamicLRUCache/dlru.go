package snowflake

type node struct {
	k, v       int
	prev, next *node
}

type LRUCache struct {
	head, tail *node
	m          map[int]*node // key -> node
	capacity   int
}

func (l *LRUCache) pushFront(n *node) {
	if l.head == nil {
		n.next, n.prev = nil, nil
		l.head, l.tail = n, n
		return
	}
	n.next, n.prev = l.head, nil
	l.head.prev = n
	l.head = n
}

func (l *LRUCache) remove(n *node) {
	if l.head == n && l.tail == n {
		l.head, l.tail = nil, nil
	} else if l.head == n {
		l.head = l.head.next
		n.next, l.head.prev = nil, nil
	} else if l.tail == n {
		l.tail = l.tail.prev
		n.prev, l.tail.next = nil, nil
	} else {
		n.prev.next, n.next.prev = n.next, n.prev
	}
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{capacity: capacity, m: make(map[int]*node, capacity+1)}
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.m[key]; ok {
		l.remove(n)
		n.v = value
		l.pushFront(n)
		return
	}
	n := &node{k: key, v: value}
	l.pushFront(n)
	l.m[key] = n

	if len(l.m) > l.capacity {
		tail := l.tail
		delete(l.m, tail.k)
		l.remove(l.tail)
	}
}

func (l *LRUCache) Get(key int) int {
	if n, ok := l.m[key]; ok {
		l.remove(n)
		l.pushFront(n)
		return n.v
	}
	return -1
}

func (l *LRUCache) Resize(newCapacity int) {
	l.capacity = newCapacity
	for len(l.m) > l.capacity {
		tail := l.tail
		delete(l.m, tail.k)
		l.remove(l.tail)
	}
}
