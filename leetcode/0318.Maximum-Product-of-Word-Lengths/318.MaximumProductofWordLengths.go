package _318_Maximum_Product_of_Word_Lengths

func maxProduct(words []string) int {
	ans, mask2len := 0, map[int]int{}
	for _, word := range words {
		var mask int
		for _, c := range word {
			mask |= 1 << int(c - 'a')
		}
		if len(word) > mask2len[mask] {
			mask2len[mask] = len(word)
		}
	}
	for x, m := range mask2len {
		for y, n := range mask2len {
			if x&y == 0 && m*n > ans {
				ans = m * n
			}
		}
	}
	return ans
}

func maxProduct2(words []string) int {
	ans, values := 0, make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			values[i] |= 1 << int(c-'a')
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if l := len(words[i]) * len(words[j]); l > ans && values[i]&values[j] == 0 {
				ans = l
			}

		}
	}
	return ans
}
