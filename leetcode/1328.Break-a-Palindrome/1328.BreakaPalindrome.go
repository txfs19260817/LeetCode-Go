package _328_Break_a_Palindrome

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	bs := []byte(palindrome)
	for i, b := range bs[:len(bs)/2] {
		if b != 'a' {
			bs[i] = 'a'
			return string(bs)
		}
	}
	bs[len(bs)-1] = 'b'
	return string(bs)
}
