package leetcode

import "container/heap"

type pt struct {
	point []int
	dist  int
}

// QuickSort
func kClosest(points [][]int, K int) [][]int {
	pts := make([]pt, len(points))
	for i, point := range points {
		pts[i] = pt{point, point[0]*point[0] + point[1]*point[1]}
	}
	qs(pts, 0, len(points)-1, K)
	for i, p := range pts[:K] {
		points[i] = p.point
	}
	return points[:K]
}

func qs(pts []pt, l, r, K int) {
	if l == r {
		return
	}
	pivot := partition(pts, l, r)
	if pivot == K {
		return
	}
	if pivot > K {
		qs(pts, l, pivot-1, K)
	} else {
		qs(pts, pivot+1, r, K)
	}
}

func partition(pts []pt, l, r int) int {
	pivotElem, i := pts[r], l
	for j := l; j < r; j++ {
		if pts[j].dist < pivotElem.dist {
			pts[j], pts[i] = pts[i], pts[j]
			i++
		}
	}
	pts[r], pts[i] = pts[i], pts[r]
	return i
}

// Heap
type ptHeap [][]int

func (h ptHeap) Len() int {
	return len(h)
}

func (h ptHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h ptHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ptHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *ptHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func kClosest2(points [][]int, K int) [][]int {
	ans, h := make([][]int, K), &ptHeap{}
	heap.Init(h)
	for _, point := range points {
		heap.Push(h, point)
		if h.Len() > K {
			heap.Pop(h)
		}
	}
	for i, p := range *h { // no need to use h.Pop() here, since we do not care about order
		ans[i] = p
	}
	return ans
}
