package leetcode

import "strconv"

type pair struct {
	max, idx int
}

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	// Get the maximum value of all numbers to the right of position i
	maxToRight := make([]pair, len(s))
	maxToRight[len(maxToRight)-1] = pair{max: num - num/10*10, idx: len(maxToRight) - 1}
	for i := len(maxToRight) - 2; i >= 0; i-- {
		if cur := int(s[i] - '0'); cur > maxToRight[i+1].max { // should be strict larger, e.g. [1]9[9]2 -> 9912
			maxToRight[i] = pair{max: cur, idx: i}
		} else {
			maxToRight[i] = maxToRight[i+1]
		}
	}
	for i := range s {
		if cur := int(s[i] - '0'); cur < maxToRight[i].max { // here is a larger number on the right side of i, do swap
			bytes := []byte(s)
			bytes[maxToRight[i].idx] = s[i] // assign the front/smaller number first
			bytes[i] = byte(maxToRight[i].max + '0')
			ans, _ := strconv.Atoi(string(bytes))
			return ans
		}
	}
	return num
}

func maximumSwap2(num int) int {
	s, lastIdx := strconv.Itoa(num), make([]int, 10)
	for i := range s {
		lastIdx[int(s[i]-'0')] = i
	}
	for i := 0; i < len(s)-1; i++ {
		for d := 9; d > int(s[i]-'0'); d-- {
			if j := lastIdx[d]; j > i {
				bytes := []byte(s)
				bytes[i], bytes[j] = bytes[j], bytes[i]
				ans, _ := strconv.Atoi(string(bytes))
				return ans
			}
		}
	}
	return num
}
