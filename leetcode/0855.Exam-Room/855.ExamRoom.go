package leetcode

import "container/heap"

type interval struct {
	l, r int
}

func (iv *interval) dist(n int) int {
	if iv.l == -1 {
		return iv.r // seat at 0, distance = r - 0
	}
	if iv.r == n {
		return (n - 1) - iv.l // seat at n-1
	}
	return (iv.r - iv.l) / 2
}

func (iv *interval) seatPos(n int) int {
	if iv.l == -1 {
		return 0
	}
	if iv.r == n {
		return n - 1
	}
	return (iv.l + iv.r) / 2
}

type intervalHeap struct {
	n int
	q []*interval
}

func (ih intervalHeap) Len() int {
	return len(ih.q)
}

func (ih intervalHeap) Less(i, j int) bool {
	inti, intj := ih.q[i], ih.q[j]
	if di, dj := inti.dist(ih.n), intj.dist(ih.n); di != dj {
		return di > dj
	}
	if pi, pj := inti.seatPos(ih.n), intj.seatPos(ih.n); pi != pj {
		return pi < pj
	}
	if inti.l != intj.l {
		return inti.l < intj.l
	}
	return inti.r < intj.r
}

func (ih intervalHeap) Swap(i, j int) {
	ih.q[i], ih.q[j] = ih.q[j], ih.q[i]
}

func (ih *intervalHeap) Push(x interface{}) {
	ih.q = append(ih.q, x.(*interval))
}

func (ih *intervalHeap) Pop() interface{} {
	n := len(ih.q)
	x := ih.q[n-1]
	ih.q = ih.q[:n-1]
	return x
}

type ExamRoom struct {
	n          int
	start, end map[int]*interval // start[l] = interval whose left endpoint is l; end[r] = interval whose right endpoint is r
	ih         intervalHeap
}

func Constructor(n int) ExamRoom {
	e := ExamRoom{n: n, start: make(map[int]*interval), end: make(map[int]*interval), ih: intervalHeap{n: n}}
	heap.Init(&e.ih)
	it := &interval{l: -1, r: n}
	e.start[it.l] = it
	e.end[it.r] = it
	heap.Push(&e.ih, it)
	return e
}

func (this *ExamRoom) add(it *interval) {
	this.start[it.l] = it
	this.end[it.r] = it
	heap.Push(&this.ih, it)
}

func (this *ExamRoom) remove(it *interval) {
	delete(this.start, it.l)
	delete(this.end, it.r)
	// lazy delete it from heap in Seat()
}

func (this *ExamRoom) Seat() int {
	var it *interval
	for {
		// pop until we get a valid interval
		it = heap.Pop(&this.ih).(*interval)
		if it == this.start[it.l] && it == this.end[it.r] {
			break
		}
	}
	seat := it.seatPos(this.n)
	this.remove(it)

	// split into (l, s) and (s, r)
	leftIt, rightIt := &interval{l: it.l, r: seat}, &interval{l: seat, r: it.r}
	this.add(leftIt)
	this.add(rightIt)
	return seat
}

func (this *ExamRoom) Leave(p int) {
	leftIt := this.end[p]    // interval ending at p
	rightIt := this.start[p] // interval starting at p
	if leftIt == nil || rightIt == nil {
		return
	}
	this.remove(leftIt)
	this.remove(rightIt)
	this.add(&interval{l: leftIt.l, r: rightIt.r})
}
