package _482_License_Key_Formatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	chars := strings.Join(strings.Split(strings.ToUpper(s), "-"), "")
	out := make([]string, 0, len(chars)/k+1)
	for i := len(chars); i >= 0; i -= k {
		start := i - k
		if start < 0 {
			start = 0
		}
		out = append(out, chars[start:i])
	}
	reverseStrings(out)
	if len(out) > 0 && len(out[0]) == 0 {
		out = out[1:]
	}
	return strings.Join(out, "-")
}

func reverseStrings(s []string) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
