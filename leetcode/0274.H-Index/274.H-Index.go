package _274_H_Index

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	bucket := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			bucket[n]++
		} else {
			bucket[c]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += bucket[i]
		if count >= i {
			return i
		}
	}
	return 0
}
