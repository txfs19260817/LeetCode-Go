package leetcode

func largestRectangleArea(heights []int) int {
	ans, monoStack := 0, make([]int, 0, len(heights)) // stores indices
	heights = append(append([]int{0}, heights...), 0)
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 0 && heights[i] < heights[monoStack[len(monoStack)-1]] {
			curHeight := heights[monoStack[len(monoStack)-1]]
			monoStack = monoStack[:len(monoStack)-1]
			width := i - monoStack[len(monoStack)-1] - 1
			if s := width * curHeight; s > ans {
				ans = s
			}
		}
		monoStack = append(monoStack, i)
	}
	return ans
}
