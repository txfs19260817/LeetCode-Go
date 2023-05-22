package leetcode

func addToArrayForm(num []int, k int) []int {
	for i := len(num) - 1; i >= 0; i-- {
		k += num[i]
		num[i] = k % 10
		k /= 10
	}
	for ; k > 9; k /= 10 {
		num = append([]int{k % 10}, num...)
	}
	if k > 0 {
		num = append([]int{k}, num...)
	}
	return num
}
