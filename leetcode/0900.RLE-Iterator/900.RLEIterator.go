package _900_RLE_Iterator

type RLEIterator struct {
	encoding   []int
	idx, count int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding, idx: -1}
}

func (this *RLEIterator) Next(n int) int {
	for n > 0 {
		for this.count == 0 {
			this.idx += 2
			if this.idx >= len(this.encoding) {
				return -1
			}
			this.count = this.encoding[this.idx-1]
		}
		if n <= this.count {
			this.count -= n
			if this.idx >= len(this.encoding) {
				return -1
			}
			return this.encoding[this.idx]
		} else {
			n -= this.count
			this.count = 0
		}
	}
	return -1
}