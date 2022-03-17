package _203_Count_Artifacts_That_Can_Be_Extracted

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	var ans int
	m := map[[2]int]bool{}
	for _, d := range dig {
		m[[2]int{d[0], d[1]}] = true
	}

	for _, artifact := range artifacts {
		top, left, bottom, right := artifact[0], artifact[1], artifact[2], artifact[3]
		blocks := make([][2]int, 0, 4)
		for x := top; x <= bottom; x++ {
			for y := left; y <= right; y++ {
				p := [2]int{x, y}
				blocks = append(blocks, p)
			}
		}
		flag := true
		for _, b := range blocks {
			if !m[b] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}
