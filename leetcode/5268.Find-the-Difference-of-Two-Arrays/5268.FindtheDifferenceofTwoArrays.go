package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	ans, m1, m2 := make([][]int, 2), map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	for v := range m1 {
		if !m2[v] {
			ans[0] = append(ans[0], v)
		}
	}
	for v := range m2 {
		if !m1[v] {
			ans[1] = append(ans[1], v)
		}
	}
	return ans
}
