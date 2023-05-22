package leetcode

func numRabbits(answers []int) int {
	var ans int
	freq := map[int]int{} // answers[i]: count
	for _, y := range answers {
		freq[y]++
	}
	for y, x := range freq {
		ans += (x + y) / (y + 1) * (y + 1)
	}
	return ans
}
