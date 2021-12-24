package in

const nodeSize = 5

type node struct {
	chars  [nodeSize]byte
	length int
	next   *node
}

type linkedList struct {
	head     *node
	totalLen int // the number of all chars in ll
}

func newLinkedList(nodes ...node) *linkedList {
	ll := &linkedList{}
	if len(nodes) > 0 {
		ll.head = &nodes[0]
		for i, p := 1, ll.head; i <= len(nodes); i++ {
			ll.totalLen += p.length
			if i < len(nodes) {
				p.next = &nodes[i]
			}
			p = p.next
		}
	}
	return ll
}

func (l linkedList) get(index int) (ans byte) {
	if index < 0 || index >= l.totalLen {
		return
	}
	for p := l.head; p != nil; p = p.next {
		if index >= p.length {
			index -= p.length
			continue
		}
		ans = p.chars[index]
		break
	}
	return
}

func (l *linkedList) insert(val byte, index int) {
	if index < 0 || index > l.totalLen { // what if index > l.totalLen, return error or set to totalLen?
		return
	}
	if l.head == nil {
		l.head = &node{chars: [nodeSize]byte{val}, length: 1}
		l.totalLen = 1
		return
	}
	for p := l.head; p != nil; p = p.next {
		if index > p.length {
			index -= p.length
			continue
		}
		if p.length == nodeSize {
			newNode := &node{next: p.next, chars: [nodeSize]byte{p.chars[nodeSize-1]}, length: 1}
			p.next = newNode
			p.length--
		}
		for i := p.length - 1; i >= index; i-- {
			p.chars[i+1] = p.chars[i]
		}
		p.chars[index] = val
		p.length++
		l.totalLen++
		break
	}
}
