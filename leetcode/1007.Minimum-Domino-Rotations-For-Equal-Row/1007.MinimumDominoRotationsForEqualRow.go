package _007_Minimum_Domino_Rotations_For_Equal_Row

func minDominoRotations(tops []int, bottoms []int) int {
	var n2is [7][]int
	ans, N := 1<<30, len(tops)
	for i := 0; i < N; i++ {
		n2is[tops[i]] = append(n2is[tops[i]], i)
		if tops[i] != bottoms[i] {
			n2is[bottoms[i]] = append(n2is[bottoms[i]], i)
		}
	}
	for n, indices := range n2is {
		if len(indices) == N {
			var t, b int
			for i := 0; i < N; i++ {
				if tops[i] == bottoms[i] {
					continue
				}
				if tops[i] == n {
					t++
				} else {
					b++
				}
			}
			ans = min(t, min(b, ans))
		}
	}
	if ans == 1<<30 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
