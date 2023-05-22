package leetcode

// Iterator Below is the interface which is already defined for you.
type Iterator struct {
	data []int
	idx  int
}

func NewIterator(data []int) *Iterator {
	return &Iterator{data, 0}
}

// hasNext returns true if the iteration has more elements.
func (i *Iterator) hasNext() bool {
	return i.idx < len(i.data)-1
}

// next returns the next element in the iteration.
func (i *Iterator) next() int {
	i.idx++
	return i.data[i.idx-1]
}

// PeekingIterator is your implementation
// https://github.com/google/guava/blob/703ef758b8621cfbab16814f01ddcc5324bdea33/guava-gwt/src-super/com/google/common/collect/super/com/google/common/collect/Iterators.java#L1125
type PeekingIterator struct {
	iter *Iterator
	// update below fields together
	nextEle   int // better to be nil
	hasPeeked bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeeked || this.iter.hasNext()
}

func (this *PeekingIterator) next() (n int) {
	if this.hasPeeked {
		n = this.nextEle
	} else {
		n = this.iter.next()
	}
	this.nextEle, this.hasPeeked = 0, false
	return
}

func (this *PeekingIterator) peek() int {
	if !this.hasPeeked {
		this.nextEle, this.hasPeeked = this.iter.next(), true
	}
	return this.nextEle
}
