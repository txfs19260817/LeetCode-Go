package _441_Arranging_Coins

import "math"

func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}
