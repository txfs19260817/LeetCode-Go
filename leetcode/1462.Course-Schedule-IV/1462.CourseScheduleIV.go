package leetcode

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	r := make([][]bool, numCourses)
	for i := range r {
		r[i] = make([]bool, numCourses)
	}
	for _, p := range prerequisites {
		r[p[0]][p[1]] = true
	}
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				r[i][j] = r[i][j] || (r[i][k] && r[k][j])
			}
		}
	}
	for i, p := range queries {
		ans[i] = r[p[0]][p[1]]
	}
	return ans
}
