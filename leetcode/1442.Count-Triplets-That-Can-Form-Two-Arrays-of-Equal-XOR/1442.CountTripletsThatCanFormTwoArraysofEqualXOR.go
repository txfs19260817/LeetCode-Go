package _442_Count_Triplets_That_Can_Form_Two_Arrays_of_Equal_XOR

func countTriplets(arr []int) int {
	ans, s, count, total := 0, 0, map[int]int{}, map[int]int{} // index:xor, index:count_i
	for k, v := range arr {
		if m, ok := count[s^v]; ok {
			ans += m*k - total[s^v]
		}
		count[s]++
		total[s] += k
		s ^= v
	}
	return ans
}

func countTriplets1(arr []int) int {
	ans, prefix := 0, make([]int, len(arr)+1)
	for i, v := range arr {
		prefix[i+1] = prefix[i] ^ v
	}
	for i := 1; i < len(arr); i++ {
		for k := i + 1; k <= len(arr); k++ {
			if prefix[i-1] == prefix[k] { // prefix[i-1]^prefix[k] == 0
				ans += k - i // k-(i-1)+1
			}
		}
	}
	return ans
}
