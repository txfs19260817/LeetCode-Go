package _867_Reverse_Prefix_of_Word

func reversePrefix(word string, ch byte) string {
	var r int
	chs := []byte(word)
	for i, c := range chs {
		if c == ch {
			r = i
			break
		}
	}
	if r == len(chs) {
		return word
	}
	for l := 0; l < r; l, r = l+1, r-1 {
		chs[l], chs[r] = chs[r], chs[l]
	}
	return string(chs)
}
