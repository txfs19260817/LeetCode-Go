package _044_Longest_Duplicate_Substring

func longestDupSubstring(s string) string {
	l, r := 1, len(s)
	for l < r {
		mid := l + (r-l)/2
		if findDuplicate(s, mid) != -1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	found := findDuplicate(s, l-1)
	if found != -1 {
		return s[found : found+l-1]
	}
	return ""
}

// Rabin-Karp
func findDuplicate(s string, length int) int {
	var h int
	for i := 0; i < length; i++ {
		h = (h*26 + int(s[i]) - 'a') % 4294967296
	}
	hash2str := map[int]string{h: s[:length]}

	// const value to be used often : 26^length % mod
	aL := 1
	for i := 1; i <= length; i++ {
		aL = (aL * 26) % 4294967296
	}

	for i := 1; i < len(s)-length+1; i++ {
		h = (h*26 - (int(s[i-1])-'a')*aL%4294967296 + 4294967296) % 4294967296
		h = (h + int(s[i+length-1]-'a')) % 4294967296
		if v, ok := hash2str[h]; ok && v == s[i:i+length] { // Be aware of hash collision!
			return i
		}
		hash2str[h] = s[i : i+length]
	}
	return -1
}
