package _370_Range_Addition

func getModifiedArray(length int, updates [][]int) []int {
	arr := make([]int, length)
	for _, update := range updates {
		start, end, inc := update[0], update[1]+1, update[2]
		arr[start] += inc
		if end < length {
			arr[end] -= inc
		}
	}
	for i := 1; i < length; i++ {
		arr[i] += arr[i-1]
	}
	return arr
}
