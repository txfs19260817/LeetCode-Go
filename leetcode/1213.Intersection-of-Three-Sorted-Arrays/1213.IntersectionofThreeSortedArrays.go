package _213_Intersection_of_Three_Sorted_Arrays

import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	ans := intersectionOfTwo(arr1, arr2)
	ans = intersectionOfTwo(ans, arr3)
	return ans
}

func intersectionOfTwo(a, b []int) []int {
	if len(a) > len(b) { // assume len(a)<len(b)
		a, b = b, a
	}
	ans := make([]int, 0, len(a))
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			ans = append(ans, a[i])
			i++
			j++
			continue
		}
		if a[i] > b[j] {
			j = sort.SearchInts(b, a[i])
		} else {
			i = sort.SearchInts(a, b[j])
		}
	}
	return ans
}

func arraysIntersection1(arr1 []int, arr2 []int, arr3 []int) []int {
	var ans []int
	m := map[int]int{}
	for _, n := range arr1 {
		m[n]++
	}
	for _, n := range arr2 {
		m[n]++
	}
	for _, n := range arr3 {
		m[n]++
	}
	for k, v := range m {
		if v == 3 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return ans
}
