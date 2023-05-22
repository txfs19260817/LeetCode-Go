package leetcode

type node struct {
	id, time int
	children []*node
}

func (root *node) dfs(ans *int, accTime int) {
	if len(root.children) == 0 { // leaf
		if accTime > *ans {
			*ans = accTime
		}
	}
	accTime += root.time
	for _, child := range root.children {
		child.dfs(ans, accTime)
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	ans, employees := 0, make([]*node, n)
	for i := 0; i < n; i++ {
		employees[i] = &node{id: i, time: informTime[i]}
	}
	for i, mId := range manager {
		if mId != -1 {
			employees[mId].children = append(employees[mId].children, employees[i])
		}
	}
	employees[headID].dfs(&ans, 0)
	return ans
}
