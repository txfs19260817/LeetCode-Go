package _841_Keys_and_Rooms

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	var dfs func(int)
	dfs = func(idx int) {
		if visited[idx] {
			return
		}
		visited[idx] = true
		for _, i := range rooms[idx] {
			dfs(i)
		}
	}
	dfs(0)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
