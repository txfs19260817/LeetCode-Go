package _290_Word_Pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	p2s, s2p, ss := map[byte]string{}, map[string]byte{}, strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}
	for i := range pattern {
		if len(p2s[pattern[i]]) > 0 && p2s[pattern[i]] != ss[i] || s2p[ss[i]] > 0 && s2p[ss[i]] != pattern[i] {
			return false
		}
		p2s[pattern[i]], s2p[ss[i]] = ss[i], pattern[i]
	}
	return true
}
