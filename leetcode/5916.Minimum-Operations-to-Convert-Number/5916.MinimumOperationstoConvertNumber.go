package _916_Minimum_Operations_to_Convert_Number

func minimumOperations(nums []int, start int, goal int) int {
	visited, q := [1001]bool{}, []int{start}
	visited[start] = true
	for step := 1; len(q) > 0; step++ {
		var p []int
		for _, top := range q {
			for _, num := range nums {
				for _, next := range []int{top + num, top - num, top ^ num} {
					if next == goal {
						return step
					}
					if next >= 0 && next <= 1000 && !visited[next] {
						visited[next] = true
						p = append(p, next)
					}
				}
			}
		}
		q = p
	}
	return -1
}
