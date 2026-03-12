package leetcode

func longestDupSubstring(s string) string {
	l, r := 1, len(s) // [1, len(s))
	for l <= r {
		mid := l + (r-l)/2
		if findDuplicate(s, mid) != -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if found := findDuplicate(s, l-1); found != -1 { // l-1 == mid
		return s[found : found+l-1]
	}
	return ""
}

// Rabin-Karp
func findDuplicate(s string, length int) int {
	const (
		base = 26
		mod  = 1 << 32
	)

	var h int
	for i := 0; i < length; i++ {
		// Interpret the substring as a base-26 number.
		h = (h*base + int(s[i]-'a')) % mod
	}
	hash2str := map[int]string{h: s[:length]}

	// basePowLen = base^length % mod, used when removing the outgoing char.
	basePowLen := 1
	for i := 0; i < length; i++ {
		basePowLen = (basePowLen * base) % mod
	}

	for i := 1; i < len(s)-length+1; i++ {
		outgoing := int(s[i-1] - 'a')
		incoming := int(s[i+length-1] - 'a')

		// Slide the window by one char:
		// 1. multiply by base to shift digits left
		// 2. subtract the outgoing char's contribution
		// 3. add mod once to keep the value non-negative before %
		h = (h*base - outgoing*basePowLen%mod + mod) % mod
		// Append the incoming char at the low digit.
		h = (h + incoming) % mod
		if v, ok := hash2str[h]; ok && v == s[i:i+length] { // be aware of hash collision!
			return i
		}
		hash2str[h] = s[i : i+length]
	}
	return -1
}
