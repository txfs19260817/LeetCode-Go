package _295_Find_Median_from_Data_Stream

import "container/heap"

type MedianFinder struct {
	small maxHeap
	large minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.small.Len() > 0 && num <= this.small[0] {
		heap.Push(&this.small, num)
		if this.small.Len() > this.large.Len()+1 {
			heap.Push(&this.large, heap.Pop(&this.small))
		}
	} else {
		heap.Push(&this.large, num)
		if this.small.Len() < this.large.Len() {
			heap.Push(&this.small, heap.Pop(&this.large))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64(this.small[0])
	}
	return (float64(this.small[0]) + float64(this.large[0])) / 2
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}
