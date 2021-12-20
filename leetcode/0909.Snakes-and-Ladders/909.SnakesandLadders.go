package _909_Snakes_and_Ladders

func snakesAndLadders(board [][]int) int {
	n := len(board)
	id2rc := func(id int) (r, c int) {
		r, c = (id-1)/n, (id-1)%n
		if r%2 == 1 {
			c = n - 1 - c
		}
		r = n - 1 - r
		return
	}
	vis := make([]bool, n*n+1)
	type pair struct{ id, step int }
	q := []pair{{1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 1; i <= 6; i++ {
			nxt := p.id + i
			if nxt > n*n { // out of range
				break
			}
			r, c := id2rc(nxt)   // id to coordinate
			if board[r][c] > 0 { // meet a snake or ladder, teleport
				nxt = board[r][c]
			}
			if nxt == n*n { // reach the goal
				return p.step + 1
			}
			if !vis[nxt] {
				vis[nxt] = true
				q = append(q, pair{nxt, p.step + 1}) // bfs any not visited next step
			}
		}
	}
	return -1
}
