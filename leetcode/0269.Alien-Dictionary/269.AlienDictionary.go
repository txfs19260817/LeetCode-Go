package leetcode

func alienOrder(words []string) string {
	// 1) collect present letters
	present, presentCount := [26]bool{}, 0
	for _, w := range words {
		for _, r := range w {
			if !present[r-'a'] {
				present[r-'a'] = true
				presentCount++
			}
		}
	}

	// 2) build graph from adjacent words
	g := [26][]int{}
	in := make([]int, 26)
	seen := make(map[[2]int]bool, 64)

	for i := 0; i+1 < len(words); i++ {
		w1, w2 := words[i], words[i+1]

		// find first different char
		j := 0
		for j < len(w1) && j < len(w2) && w1[j] == w2[j] {
			j++
		}

		// prefix invalid: "abc" before "ab"
		if j == len(w2) && len(w1) > len(w2) {
			return ""
		}

		// add edge if there is a differing char
		if j < len(w1) && j < len(w2) {
			u := int(w1[j] - 'a')
			v := int(w2[j] - 'a')
			e := [2]int{u, v}
			if !seen[e] {
				seen[e] = true
				g[u] = append(g[u], v)
				in[v]++
			}
		}
	}

	// 3) Kahn topological sort over present letters
	q := make([]int, 0, presentCount)
	for c, p := range present {
		if p && in[c] == 0 {
			q = append(q, c)
		}
	}

	ans := make([]byte, 0, presentCount)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		ans = append(ans, byte(u+'a'))
		for _, v := range g[u] {
			if in[v]--; in[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(ans) != presentCount {
		return ""
	}
	return string(ans)
}
