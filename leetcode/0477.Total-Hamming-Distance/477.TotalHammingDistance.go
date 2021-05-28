package _477_Total_Hamming_Distance

func totalHammingDistance(nums []int) int {
	var ans int
	for i := 0; i < 30; i++ {
		var count int
		for _, num := range nums {
			count += num >> i & 1
		}
		ans += count * (len(nums) - count)
	}
	return ans
}
