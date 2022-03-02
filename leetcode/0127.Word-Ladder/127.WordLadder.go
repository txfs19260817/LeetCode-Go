package _127_Word_Ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	q := []string{beginWord}
	for step := 1; len(q) > 0; step++ {
		var p []string
		for _, s := range q {
			for i := 0; i < len(s); i++ {
				bs := []byte(s)
				for j := 0; j < 26; j++ {
					bs[i] = byte('a' + j)
					newS := string(bs)
					if !wordSet[newS] {
						continue
					}
					if newS == endWord {
						return step + 1
					}
					delete(wordSet, newS)
					p = append(p, newS)
				}
			}
		}
		q = p
	}
	return 0
}
