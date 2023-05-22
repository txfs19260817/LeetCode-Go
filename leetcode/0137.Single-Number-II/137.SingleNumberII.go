package leetcode

import "math"

func singleNumber(nums []int) int {
	var X, Y int
	for _, n := range nums {
		Y = Y ^ n & ^X
		X = X ^ n & ^Y
	}
	return Y
}

func singleNumber1(nums []int) int {
	ans, pos, neg := 0, make([]int, 32), make([]int, 32)
	for _, num := range nums {
		if num >= 0 {
			setArray(num, &pos)
		} else {
			setArray(-num, &neg)
		}
	}
	for i, num := range pos {
		ans += (num % 3) * int(math.Pow(2, float64(i)))
	}
	for i, num := range neg {
		ans -= (num % 3) * int(math.Pow(2, float64(i)))
	}
	return ans
}

func setArray(num int, arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		if num&1 != 0 {
			(*arr)[i]++
		}
		num >>= 1
	}
}
