package _763_Partition_Labels

func partitionLabels(S string) []int {
	var ans []int
	last := map[byte]int{}
	for i, c := range S {
		last[byte(c)] = i
	}
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = last[S[start]]
		for i := start; i <= end; i++ {
			if end < last[S[i]] {
				end = last[S[i]]
			}
		}
		ans = append(ans, end-start+1)
	}
	return ans
}
