package leetcode

import "math/rand"

type Solution struct {
	proj        map[int]int
	m, n, total int
}

func Constructor(m int, n int) Solution {
	return Solution{proj: make(map[int]int), m: m, n: n, total: m * n}
}

func (this *Solution) Flip() []int {
	var ans []int
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.proj[x]; ok {
		ans = []int{y / this.n, y % this.n}
	} else {
		ans = []int{x / this.n, x % this.n}
	}

	if y, ok := this.proj[this.total]; ok {
		this.proj[x] = y
	} else {
		this.proj[x] = this.total
	}
	return ans
}

func (this *Solution) Reset() {
	this.proj, this.total = make(map[int]int), this.m*this.n
}
