package k

import (
	"bytes"
	"strings"
)

func reflowAndJustify(lines []string, maxLen int) (ans []string) {
	words, curLen := strings.Split(strings.Join(lines, " "), " "), 0
	for i := 0; i < len(words); i++ {
		var count int
		start := i
		for ; i < len(words) && curLen+len(words[i]) <= maxLen; i++ {
			count++
			curLen += len(words[i]) + 1
		}
		i--
		curLen-- // extra space
		if count == 1 {
			ans = append(ans, words[i])
			break
		}
		curLine := words[start : start+count]
		filler := maxLen - curLen + count - 1
		avg, mod := filler/(count-1), filler%(count-1)
		var s string
		for j, w := range curLine {
			if j == len(curLine)-1 {
				s += w
				break
			}
			s += w + string(bytes.Repeat([]byte{'-'}, avg))
			if mod > 0 {
				s += "-"
				mod--
			}
		}
		ans = append(ans, s)
		curLen = 0
	}
	return
}
