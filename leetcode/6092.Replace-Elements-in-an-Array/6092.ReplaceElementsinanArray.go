package leetcode

func arrayChange(nums []int, operations [][]int) []int {
	n2i := map[int]int{}
	for i, num := range nums {
		n2i[num] = i
	}
	for _, op := range operations {
		if j, ok := n2i[op[0]]; ok {
			n2i[op[1]] = j
			delete(n2i, op[0])
			nums[j] = op[1]
		}
	}
	return nums
}
