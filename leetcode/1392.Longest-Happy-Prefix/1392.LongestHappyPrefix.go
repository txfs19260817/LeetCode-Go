package leetcode

func longestPrefix(s string) string {
	h1, h2, length, mul, MOD := 0, 0, 0, 1, 1000000007
	for l, r := 0, len(s)-1; r > 0; l, r = l+1, r-1 {
		h1 = (h1*26 + int(s[l]-'a')) % MOD
		h2 = (h2 + int(s[r]-'a')*mul) % MOD
		mul = mul * 26 % MOD
		if h1 == h2 {
			//if s[:l+1-length] == s[r:len(s)-length] && s[length:l+1] == s[r+length:] { // handle collisions
			length = l + 1
			//}
		}
	}
	return s[:length]
}
