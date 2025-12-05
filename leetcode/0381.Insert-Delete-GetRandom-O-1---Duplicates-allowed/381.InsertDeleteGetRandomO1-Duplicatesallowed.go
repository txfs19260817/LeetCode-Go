package leetcode

import (
	"math/rand"
)

type RandomizedCollection struct {
	nums      []int
	n2indices map[int]map[int]bool
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{n2indices: map[int]map[int]bool{}}
}

func (r *RandomizedCollection) Insert(val int) bool {
	ok := len(r.n2indices[val]) > 0
	if !ok {
		r.n2indices[val] = map[int]bool{}
	}
	r.nums = append(r.nums, val)
	r.n2indices[val][len(r.nums)-1] = true

	return !ok
}

func (r *RandomizedCollection) Remove(val int) bool {
	if len(r.n2indices[val]) == 0 {
		return false
	}
	iVal := -1
	for iVal = range r.n2indices[val] {
		break
	}

	iLast := len(r.nums) - 1
	lastVal := r.nums[iLast]

	r.nums[iVal] = lastVal
	r.nums = r.nums[:iLast]

	delete(r.n2indices[val], iVal)
	delete(r.n2indices[lastVal], iLast)
	if iVal != iLast {
		r.n2indices[lastVal][iVal] = true
	}
	return true
}

func (r *RandomizedCollection) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
