package _541_Reverse_String_II

func reverseStr(s string, k int) string {
	bs := []byte(s)
	for l := 0; l < len(s); l += 2 * k {
		r := l + k
		if r > len(s) {
			r = len(s)
		}
		for i, j := l, r-1; i < j; i, j = i+1, j-1 {
			bs[i], bs[j] = bs[j], bs[i]
		}
	}
	return string(bs)
}
