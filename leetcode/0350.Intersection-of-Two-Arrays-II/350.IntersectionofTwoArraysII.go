package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, n := range nums1 {
		m[n]++
	}
	res := make([]int, 0)
	for _, n := range nums2 {
		if m[n] > 0 {
			m[n]--
			res = append(res, n)
		}
	}
	return res
}
