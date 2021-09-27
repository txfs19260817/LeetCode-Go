package _917_Reverse_Only_Letters

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	for l, r := 0, len(b)-1; l < r; {
		if !(b[l] >= 'a' && b[l] <= 'z') && !(b[l] >= 'A' && b[l] <= 'Z') {
			l++
			continue
		}
		if !(b[r] >= 'a' && b[r] <= 'z') && !(b[r] >= 'A' && b[r] <= 'Z') {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l, r = l+1, r-1
	}
	return string(b)
}
