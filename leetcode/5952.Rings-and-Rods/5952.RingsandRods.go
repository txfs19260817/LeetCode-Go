package _952_Rings_and_Rods

func countPoints(rings string) int {
	ans, m := 0, map[int]map[byte]bool{}
	for i := 0; i < len(rings); i += 2 {
		idx := int(rings[i+1] - '0')
		if m[idx] == nil {
			m[idx] = map[byte]bool{rings[i]: true}
		} else {
			m[idx][rings[i]] = true
		}
	}
	for _, set := range m {
		if len(set) == 3 {
			ans++
		}
	}
	return ans
}
