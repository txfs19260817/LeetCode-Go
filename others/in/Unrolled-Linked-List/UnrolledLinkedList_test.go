package in

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnrolledLinkedList(t *testing.T) {
	nodeList := []node{
		{chars: [nodeSize]byte{'a', 'b'}, length: 2},
		{chars: [nodeSize]byte{'b'}, length: 1},
		{chars: [nodeSize]byte{'a', 'b', 'c', 'd', 'e'}, length: 5},
	}
	ll := newLinkedList(nodeList...)
	assert.Equal(t, 8, ll.totalLen)
	assert.Equal(t, byte('a'), ll.get(0))
	assert.Equal(t, byte('b'), ll.get(2))
	assert.Equal(t, byte('a'), ll.get(3))
	ll.insert('c', 2)
	assert.Equal(t, [nodeSize]byte{'a', 'b', 'c'}, ll.head.chars)
	ll.insert('f', 4)
	ll.insert('e', 4)
	ll.insert('d', 4)
	ll.insert('c', 4)
	assert.Equal(t, [nodeSize]byte{'b', 'c', 'd', 'e', 'f'}, ll.head.next.chars)
	ll.insert('x', 4)
	assert.Equal(t, [nodeSize]byte{'b', 'x', 'c', 'd', 'e'}, ll.head.next.chars)
	assert.Equal(t, [nodeSize]byte{'f'}, ll.head.next.next.chars)
	assert.Equal(t, [nodeSize]byte{'a', 'b', 'c', 'd', 'e'}, ll.head.next.next.next.chars)
	assert.Equal(t, 14, ll.totalLen)
}
