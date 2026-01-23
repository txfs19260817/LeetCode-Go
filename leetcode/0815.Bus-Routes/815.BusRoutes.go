package leetcode

func numBusesToDestination(routes [][]int, source int, target int) int {
	stop2buses := map[int][]int{}
	for i, route := range routes {
		for _, stop := range route {
			stop2buses[stop] = append(stop2buses[stop], i)
		}
	}

	if len(stop2buses[source]) == 0 || len(stop2buses[target]) == 0 {
		if source == target {
			return 0
		}
		return -1
	}

	dist := map[int]int{source: 0}
	q := []int{source}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		distX := dist[x]
		for _, bus := range stop2buses[x] {
			for _, y := range routes[bus] {
				if _, ok := dist[y]; !ok {
					dist[y] = distX + 1
					q = append(q, y)
				}
			}
			routes[bus] = nil
		}
	}

	if ans, ok := dist[target]; ok {
		return ans
	}
	return -1
}
