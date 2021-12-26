package _965_Intervals_Between_Identical_Elements

func getDistances(arr []int) []int64 {
	ans, num2i := make([]int64, len(arr)), map[int][]int{}
	for i, num := range arr {
		num2i[num] = append(num2i[num], i)
	}
	for _, indices := range num2i {
		prefix := append([]int{}, indices...)
		for i := 1; i < len(prefix); i++ {
			prefix[i] += prefix[i-1]
		}
		for i, index := range indices {
			ans[index] += int64((i+1)*index-prefix[i]) + int64(prefix[len(prefix)-1]-prefix[i]-(len(indices)-i-1)*index)
		}
	}
	return ans
}
