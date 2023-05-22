package leetcode

func findRepeatedDnaSequences(s string) []string {
	var ans []string
	bitSet, dict, L, bitmask := map[int]int{}, map[byte]int{'A': 0, 'T': 1, 'C': 2, 'G': 3}, 10, 0
	for i := 0; i < len(s); i++ {
		bitmask = bitmask<<2 | dict[s[i]]
		bitmask &= (1<<32 - 1) >> (32 - L*2) // 1111 1111 1111 1111 1111
		if i >= L-1 {                        // start from s[0:10)
			if bitSet[bitmask] == 1 {
				ans = append(ans, s[i-L+1:i+1])
			}
			bitSet[bitmask]++
		}
	}
	return ans
}
