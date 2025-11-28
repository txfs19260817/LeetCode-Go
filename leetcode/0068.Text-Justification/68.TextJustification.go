package leetcode

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var cnt, lastIdx int
	var container [][]string
	container = append(container, make([]string, 0))
	addZero := func(s string) string {
		return s + strings.Repeat(" ", maxWidth-len(s))
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
					s += strings.Repeat(" ", spaces) + container[len(container)-1][i]
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

func fullJustify2(words []string, maxWidth int) []string {
	var ans []string
	for i := 0; i < len(words); {
		var curLine []string
		var cnt int // count all characters in words[i] and spaces between words
		for ; i < len(words) && cnt+1+len(words[i]) <= maxWidth; i++ {
			curLine = append(curLine, words[i])
			cnt += 1 + len(words[i])
		}
		if i < len(words) && cnt+len(words[i]) == maxWidth { // last word is an exact fit
			curLine = append(curLine, words[i])
			cnt += len(words[i])
			i++
		} else {
			cnt-- // last word is not an exact fit, so we need to remove the last space
		}

		if len(curLine) == 1 || i == len(words) { // only one word or last word, need to meet the requirement of last line
			joined := strings.Join(curLine, " ")
			ans = append(ans, joined+strings.Repeat(" ", maxWidth-len(joined)))
			continue
		}

		intervals := len(curLine) - 1              // number of intervals between words
		leftSpaces := maxWidth - cnt               // number of spaces to be distributed
		spacesEach := leftSpaces/intervals + 1     // number of spaces between words, plus one for we've counted it in cnt
		additionalSpaces := leftSpaces % intervals // number of additional spaces to be distributed for the first intervals

		var sb strings.Builder
		sb.Grow(maxWidth)
		for j, w := range curLine {
			sb.WriteString(w)       // write the word
			if j < len(curLine)-1 { // not the last word
				nSpaces := spacesEach // number of spaces between words
				if additionalSpaces > 0 {
					nSpaces++ // plus one for the additional space
					additionalSpaces--
				}
				sb.WriteString(strings.Repeat(" ", nSpaces)) // write the spaces
			}
		}
		ans = append(ans, sb.String())
	}
	return ans
}
