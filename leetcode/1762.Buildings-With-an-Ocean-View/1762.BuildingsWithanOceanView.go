package leetcode

func findBuildings(heights []int) []int {
	var idxStack []int
	for i, height := range heights {
		for len(idxStack) > 0 && heights[idxStack[len(idxStack)-1]] <= height {
			idxStack = idxStack[:len(idxStack)-1]
		}
		idxStack = append(idxStack, i)
	}
	return idxStack
}
