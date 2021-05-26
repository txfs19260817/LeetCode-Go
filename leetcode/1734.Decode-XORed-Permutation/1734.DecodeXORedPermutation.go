package _734_Decode_XORed_Permutation

func decode(encoded []int) []int {
	var total, odd int
	for i := 1; i <= len(encoded)+1; i++ {
		total ^= i
	}
	for i := 1; i < len(encoded); i += 2 {
		odd ^= encoded[i]
	}
	perm := make([]int, len(encoded)+1)
	perm[0] = total ^ odd
	for i := 1; i < len(perm); i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}
	return perm
}
