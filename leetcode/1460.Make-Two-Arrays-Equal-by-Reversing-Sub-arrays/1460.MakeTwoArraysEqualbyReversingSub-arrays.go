package leetcode

import "sort"

func canBeEqual(target []int, arr []int) bool {
	var m [1001]int
	for i := 0; i < len(target); i++ {
		m[target[i]]++
		m[arr[i]]--
	}
	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}

func canBeEqual1(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
