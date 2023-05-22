package leetcode

import (
	"math"
	"math/rand"
)

type RandomizedCollection struct {
	v2indices map[int]map[int]bool
	values    []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{v2indices: make(map[int]map[int]bool)}
}

func (this *RandomizedCollection) Insert(val int) bool {
	ans := true
	if _, ok := this.v2indices[val]; ok {
		ans = false
	}
	this.values = append(this.values, val)
	if this.v2indices[val] == nil {
		this.v2indices[val] = make(map[int]bool)
	}
	this.v2indices[val][len(this.values)-1] = true
	return ans
}

func (this *RandomizedCollection) Remove(val int) bool {
	indicesSet, ok := this.v2indices[val]
	if !ok || len(indicesSet) == 0 {
		return false
	}
	// get & remove an index of val from indices set
	var valIdx int
	for i := range indicesSet {
		valIdx = i
		break
	}
	delete(indicesSet, valIdx)
	if len(indicesSet) == 0 {
		delete(this.v2indices, val)
	} else {
		this.v2indices[val] = indicesSet
	}

	this.values[valIdx] = this.values[len(this.values)-1] // replace target val with the last element
	// add a new index for the last element
	if this.v2indices[this.values[len(this.values)-1]] == nil {
		this.v2indices[this.values[len(this.values)-1]] = make(map[int]bool)
	}
	this.v2indices[this.values[len(this.values)-1]][valIdx] = true
	// remove the old one
	delete(this.v2indices[this.values[len(this.values)-1]], len(this.values)-1)
	if len(this.v2indices[this.values[len(this.values)-1]]) == 0 {
		delete(this.v2indices, this.values[len(this.values)-1])
	}
	this.values = this.values[:len(this.values)-1] // pop

	return true
}

func (this *RandomizedCollection) GetRandom() int {
	return this.values[int(math.Floor(rand.Float64()*float64(len(this.values))))]
}
