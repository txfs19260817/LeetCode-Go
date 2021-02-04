package _923_3Sum_With_Multiplicity

import "sort"

func threeSumMulti(arr []int, target int) int {
	ans, freq := 0, map[int]int{}
	for _, a := range arr {
		freq[a]++
	}

	unique := make([]int, 0, len(freq))
	for k := range freq {
		unique = append(unique, k)
	}
	sort.Ints(unique)

	for i := 0; i < len(unique); i++ {
		ni := freq[unique[i]]
		if ni >= 3 && unique[i]*3 == target {
			ans += ni * (ni - 1) * (ni - 2) / 6
		}
		for j := i + 1; j < len(unique); j++ {
			nj := freq[unique[j]]
			if ni > 1 && unique[i]*2+unique[j] == target {
				ans += nj * ni * (ni - 1) / 2
			}
			if nj > 1 && unique[j]*2+unique[i] == target {
				ans += ni * nj * (nj - 1) / 2
			}
			if r := target - unique[i] - unique[j]; freq[r] > 0 && r > unique[j] {
				ans += ni * nj * freq[r]
			}
		}
	}
	return ans % 1000000007
}
