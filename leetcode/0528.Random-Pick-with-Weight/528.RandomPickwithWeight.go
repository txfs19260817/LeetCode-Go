package _528_Random_Pick_with_Weight

import (
	"math/rand"
	"sort"
)

type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{pre: w}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, r)
}
