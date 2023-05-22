package leetcode

func deleteAndEarn(nums []int) int {
	freq := make([]int, 10001)
	for _, num := range nums {
		freq[num]++
	}
	first, second := 0, freq[1] // freq[0]*0, freq[1]*1
	for i := 2; i < len(freq); i++ {
		first, second = second, max(second, first+freq[i]*i)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
