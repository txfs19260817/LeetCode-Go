package _319_Bulb_Switcher

import "math"

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + .5)) // there are floor(sqrt(n)) numbers has odd amount of factors
}
