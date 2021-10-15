package _000_Reverse_Prefix_of_Word

func reversePrefix(word string, ch byte) string {
	end := -1
	for i, r := range word {
		if byte(r) == ch {
			end = i
			break
		}
	}
	if end == -1 {
		return word
	}
	bytes := []byte(word)
	for start := 0; start < end; start, end = start+1, end-1 {
		bytes[start], bytes[end] = bytes[end], bytes[start]
	}
	return string(bytes)
}
