package _969_Destroying_Asteroids

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for len(asteroids) > 0 {
		idx := sort.SearchInts(asteroids, mass)
		for idx < len(asteroids) && asteroids[idx] == mass {
			idx++
		}
		if idx == 0 {
			return false
		}
		if idx == len(asteroids) {
			break
		}
		for i := 0; i < idx; i++ {
			mass += asteroids[i]
		}
		asteroids = asteroids[idx:]
	}
	return true
}
