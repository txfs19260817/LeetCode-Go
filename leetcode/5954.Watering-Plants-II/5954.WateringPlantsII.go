package leetcode

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var ans int
	for a, b, ca, cb := 0, len(plants)-1, capacityA, capacityB; a <= b; a, b = a+1, b-1 {
		if a == b {
			if ca >= cb {
				if ca < plants[a] {
					ans++
				}
			} else if cb < plants[b] {
				ans++
			}
			break
		}
		if ca < plants[a] {
			ans++
			ca = capacityA
		}
		ca -= plants[a]
		if cb < plants[b] {
			ans++
			cb = capacityB
		}
		cb -= plants[b]
	}
	return ans
}
