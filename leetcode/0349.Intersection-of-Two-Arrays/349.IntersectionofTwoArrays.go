package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, n := range nums1 {
		m[n] = true
	}
	res := make([]int, 0)
	for _, n := range nums2 {
		if m[n] {
			res = append(res, n)
			m[n] = false
		}
	}
	return res
}
