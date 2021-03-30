package _470_Implement_Rand10_Using_Rand7

import "math/rand"

func rand10() int {
	t := 50
	for t > 39 {
		t = (rand7()-1)*7 + (rand7() - 1)
	}
	return t%10 + 1
}

func rand10_1() int {
	for {
		r, c := rand7(), rand7()
		if num := (r-1)*7 + c; num <= 40 {
			return num%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
