package _583_Count_Unhappy_Friends

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	ans, rank, match := 0, make([][]int, n), make([]int, n)
	for i, preference := range preferences {
		rank[i] = make([]int, n)
		for j, p := range preference {
			rank[i][p] = j
		}
	}
	for _, p := range pairs {
		match[p[0]], match[p[1]] = p[1], p[0]
	}
	for x, y := range match {
		idx := rank[x][y]
		for _, u := range preferences[x][:idx] {
			if v := match[u]; rank[u][x] < rank[u][v] {
				ans++
				break
			}
		}
	}
	return ans
}
