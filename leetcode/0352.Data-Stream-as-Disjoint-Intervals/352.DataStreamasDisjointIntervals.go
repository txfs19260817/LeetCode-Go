package leetcode

type SummaryRanges struct {
	intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{intervals: [][]int{}}
}

func (this *SummaryRanges) AddNum(val int) {
	idx := this.Search(val)         // the index to insert val
	if idx >= len(this.intervals) { // larger than all values in this.intervals
		if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] { // the last element is [x, val-1], val => [x, val]
			this.intervals[idx-1][1]++
		} else {
			this.intervals = append(this.intervals, []int{val, val}) // append [val, val]
		}
		return
	}
	if this.intervals[idx][0] <= val && val <= this.intervals[idx][1] { // val is in this.intervals[idx]
		return
	}
	if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] && val+1 == this.intervals[idx][0] { // val connects two adjacent intervals. i.e. [x, val-1], val, [val+1, y] => [x, y]
		this.intervals = append(append(this.intervals[:idx-1], []int{this.intervals[idx-1][0], this.intervals[idx][1]}), this.intervals[idx+1:]...)
		return
	}
	if val+1 == this.intervals[idx][0] { // only val, [val+1, y] => [val, y]
		this.intervals[idx][0]--
		return
	}
	if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] { // only [x, val-1], val => [x, val]
		this.intervals[idx-1][1]++
		return
	}
	if idx == 0 && val < this.intervals[idx][0] { // smaller than all values in this.intervals
		this.intervals = append([][]int{{val, val}}, this.intervals...) // insert it to the first place
		return
	}
	if idx-1 >= 0 && val-1 != this.intervals[idx-1][1] && val+1 != this.intervals[idx][0] { // insert [val, val] to the `idx` place of this.intervals
		tmp := make([][]int, len(this.intervals[:idx]))
		copy(tmp, this.intervals[:idx])
		tmp = append(tmp, []int{val, val})
		this.intervals = append(tmp, this.intervals[idx:]...)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.intervals
}

// Search should perform like sort.Search,
// i.e. the return value is the index to insert val
func (this *SummaryRanges) Search(val int) int {
	// tell caller to append val to this.intervals when it is an empty slice
	if len(this.intervals) == 0 {
		return 0
	}
	// binary search
	l, r := 0, len(this.intervals)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if interval := this.intervals[mid]; val < interval[0] {
			r = mid
		} else if val > interval[1] {
			l = mid
		} else { // val has been already included in an existed interval
			return mid
		}
	}
	// e.g. intervals = [[1, 2], [5, 6]] => l = 0, r = 1
	switch {
	case val > this.intervals[r][1]: // val > 6 => idx = 2 (append)
		return r + 1
	case this.intervals[l][1] < val && val <= this.intervals[r][1]: // val in (2, 6] => idx = 1
		return r
	default: // val <= 2 => 0
		return l
	}
}
