package leetcode

import "sort"

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}

	cars := make([]car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = car{position[i], float64(target-position[i]) / float64(speed[i])}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleet, lastTime := 0, 0.0
	for _, c := range cars {
		if c.time > lastTime {
			lastTime = c.time
			fleet++
		}
	}
	return fleet
}

type car struct {
	position int
	time     float64
}
