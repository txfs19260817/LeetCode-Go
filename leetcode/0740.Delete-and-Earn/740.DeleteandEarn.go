package _740_Delete_and_Earn

func deleteAndEarn(nums []int) int {
	var maxV int
	for _, num := range nums {
		maxV = max(maxV, num)
	}
	freq := make([]int, maxV+1)
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
