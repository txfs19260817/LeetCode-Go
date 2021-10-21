package _380_Insert_Delete_GetRandom_O_1

import (
	"math"
	"math/rand"
)

type RandomizedSet struct {
	k2idx map[int]int
	keys  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{k2idx: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.k2idx[val]; ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.k2idx[val] = len(this.keys) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.k2idx[val]
	if !ok {
		return false
	}
	this.keys[idx] = this.keys[len(this.keys)-1]
	this.keys = this.keys[:len(this.keys)-1]
	delete(this.k2idx, val)
	if idx < len(this.keys) {
		this.k2idx[this.keys[idx]] = idx
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.keys[int(math.Floor(rand.Float64()*float64(len(this.keys))))]
}
