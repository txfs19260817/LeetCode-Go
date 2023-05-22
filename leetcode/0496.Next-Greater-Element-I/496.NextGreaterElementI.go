package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans, monoStack, nextGreater := make([]int, len(nums1)), make([]int, 0, len(nums2)), map[int]int{}
	for _, n := range nums2 {
		for len(monoStack) > 0 && n > monoStack[len(monoStack)-1] {
			nextGreater[monoStack[len(monoStack)-1]] = n
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, n)
	}
	for len(monoStack) > 0 {
		nextGreater[monoStack[len(monoStack)-1]] = -1
		monoStack = monoStack[:len(monoStack)-1]
	}
	for i, n := range nums1 {
		ans[i] = nextGreater[n]
	}
	return ans
}
