package leetcode

func findKthPositive(arr []int, k int) int {
	var i, p int
	for i = 1; k != 0; i++ {
		if p < len(arr) && i == arr[p] {
			p++
			continue
		}
		k--
	}
	return i - 1
}
