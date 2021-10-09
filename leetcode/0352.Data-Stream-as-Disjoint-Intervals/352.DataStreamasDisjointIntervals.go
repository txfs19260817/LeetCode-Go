package _352_Data_Stream_as_Disjoint_Intervals

type SummaryRanges struct {
	intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{intervals: [][]int{}}
}

func (this *SummaryRanges) AddNum(val int) {
	idx := this.Search(val)
	if idx >= len(this.intervals) {
		if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] {
			this.intervals[idx-1][1]++
		} else {
			this.intervals = append(this.intervals, []int{val, val})
		}
		return
	}
	if this.intervals[idx][0] <= val && val <= this.intervals[idx][1] {
		return
	}
	if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] && val+1 == this.intervals[idx][0] {
		this.intervals = append(append(this.intervals[:idx-1], []int{this.intervals[idx-1][0], this.intervals[idx][1]}), this.intervals[idx+1:]...)
		return
	}
	if val+1 == this.intervals[idx][0] {
		this.intervals[idx][0]--
		return
	}
	if idx-1 >= 0 && val-1 == this.intervals[idx-1][1] {
		this.intervals[idx-1][1]++
		return
	}
	if idx == 0 && val < this.intervals[idx][0] {
		this.intervals = append([][]int{{val, val}}, this.intervals...)
		return
	}
	if idx-1 >= 0 && val-1 != this.intervals[idx-1][1] && val+1 != this.intervals[idx][0] {
		tmp := make([][]int, len(this.intervals[:idx]))
		copy(tmp, this.intervals[:idx])
		tmp = append(tmp, []int{val, val})
		this.intervals = append(tmp, this.intervals[idx:]...)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.intervals
}

func (this *SummaryRanges) Search(val int) int {
	if len(this.intervals) == 0 {
		return 0
	}
	l, r := 0, len(this.intervals)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if interval := this.intervals[mid]; val < interval[0] {
			r = mid
		} else if val > interval[1] {
			l = mid
		} else {
			return mid
		}
	}
	switch {
	case val > this.intervals[r][1]:
		return r + 1
	case this.intervals[l][1] < val && val <= this.intervals[r][1]:
		return r
	default:
		return l
	}
}
