package _318_Maximum_Product_of_Word_Lengths

func maxProduct(words []string) int {
	if len(words) == 0 {
		return 0
	}
	ans, values := 0, make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			values[i] |= 1 << int(c-'a')
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if values[i]&values[j] == 0 {
				if l := len(words[i]) * len(words[j]); l > ans {
					ans = l
				}
			}
		}
	}
	return ans
}
