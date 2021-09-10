package _068_Text_Justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var cnt, lastIdx int
	var container [][]string
	container = append(container, make([]string, 0))
	spacesGen := func(n int) string {
		spaces := make([]byte, n)
		for i := range spaces {
			spaces[i] = ' '
		}
		return string(spaces)
	}
	addZero := func(s string) string {
		return s + spacesGen(maxWidth-len(s))
	}
	for no, word := range words {
		cnt += len(word)
		if cnt > maxWidth {
			cnt -= len(word) + len(container[len(container)-1]) // char count
			s := container[len(container)-1][0]
			if len(container[len(container)-1]) == 1 {
				s = addZero(s)
			} else {
				d, r := (maxWidth-cnt)/(len(container[len(container)-1])-1), (maxWidth-cnt)%(len(container[len(container)-1])-1)
				for i := 1; i < len(container[len(container)-1]); i++ {
					spaces := d
					if r > 0 {
						spaces, r = spaces+1, r-1
					}
					s += spacesGen(spaces) + container[len(container)-1][i]
				}
			}
			ans = append(ans, s)
			lastIdx = no
			cnt = len(word)
			container = append(container, make([]string, 0))
		}
		cnt++
		container[len(container)-1] = append(container[len(container)-1], word)
	}
	ans = append(ans, addZero(strings.Join(words[lastIdx:], " ")))
	return ans
}
