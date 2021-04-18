package _217_Contains_Duplicate

func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
