package leetcode

func findWords(board [][]byte, words []string) []string {
	dir := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	ans, trie := make([]string, 0, len(words)), Trie{}
	for _, word := range words {
		trie.Insert(word)
	}
	var dfs func(i, j int, p *Trie)
	dfs = func(i, j int, p *Trie) {
		c := board[i][j]
		if c == '#' || p.children[c-'a'] == nil {
			return
		}
		p = p.children[c-'a']
		if len(p.word) > 0 { // de-dup
			ans = append(ans, p.word)
			p.word = ""
		}
		board[i][j] = '#'
		for _, d := range dir {
			if x, y := i+d[0], j+d[1]; x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
				dfs(x, y, p)
			}
		}
		board[i][j] = c
	}
	for i, line := range board {
		for j := range line {
			dfs(i, j, &trie)
		}
	}
	return ans
}

type Trie struct {
	children [26]*Trie
	word     string
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	p := t
	for _, ch := range word {
		ch -= 'a'
		if p.children[ch] == nil {
			p.children[ch] = &Trie{}
		}
		p = p.children[ch]
	}
	p.word = word
}
