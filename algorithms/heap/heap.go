package main

import "fmt"

type heap struct {
	array     []int
	size, cap int
}

func newHeap(n int) *heap {
	return &heap{cap: n, array: make([]int, 0, n)}
}

func (h heap) leaf(index int) bool {
	return index >= (h.size/2) && index <= h.size
}

func (h heap) parent(index int) int {
	return (index - 1) / 2
}

func (h heap) leftChild(index int) int {
	return 2*index + 1
}

func (h heap) rightChild(index int) int {
	return 2*index + 2
}

func (h *heap) push(item int) {
	if h.size >= h.cap {
		h.pop()
	}
	h.array = append(h.array, item)
	h.size++
	h.upHeapify(h.size - 1)
}

func (h heap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h heap) less(i, j int) bool {
	return h.array[i] < h.array[j]
}

func (h *heap) upHeapify(index int) {
	for h.less(index, h.parent(index)) {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *heap) downHeapify(current int) {
	if h.leaf(current) {
		return
	}
	top, leftChildIndex, rightRightIndex := current, h.leftChild(current), h.rightChild(current)
	if leftChildIndex < h.size && h.less(leftChildIndex, top) {
		top = leftChildIndex
	}
	if rightRightIndex < h.size && h.less(rightRightIndex, top) {
		top = rightRightIndex
	}
	if top != current {
		h.swap(current, top)
		h.downHeapify(top)
	}
}

func (h *heap) fix() {
	for index := (h.size / 2) - 1; index >= 0; index-- {
		h.downHeapify(index)
	}
}

func (h *heap) pop() int {
	top := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:(h.size)-1]
	h.size--
	h.downHeapify(0)
	return top
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := newHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		minHeap.push(inputArray[i])
	}
	//minHeap.fix()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.pop())
	}
}
