package leetcode

const N = 300

type HitCounter struct {
	ts   [N]int
	hits [N]int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	i := timestamp % N
	if this.ts[i] == timestamp {
		this.hits[i]++
	} else {
		this.ts[i] = timestamp
		this.hits[i] = 1
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	var ans int
	for i := range N {
		if timestamp-this.ts[i] < N {
			ans += this.hits[i]
		}
	}
	return ans
}
