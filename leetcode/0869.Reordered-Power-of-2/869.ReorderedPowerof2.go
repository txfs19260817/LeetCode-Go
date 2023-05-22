package leetcode

import (
	"sort"
	"strconv"
)

func reorderedPowerOf2(n int) bool {
	countDigits := func(n int) (cnt [10]int) {
		for n > 0 {
			cnt[n%10]++
			n /= 10
		}
		return
	}
	powerOf2Digits := map[[10]int]bool{}
	for i := 1; i <= 1e9; i <<= 1 {
		powerOf2Digits[countDigits(i)] = true
	}

	return powerOf2Digits[countDigits(n)]
}

func reorderedPowerOf22(n int) bool {
	// get the upper and lower bounds of any possible reordered n
	digits := []byte(strconv.Itoa(n))
	sort.Slice(digits, func(i, j int) bool { return digits[i] > digits[j] })
	upper, _ := strconv.Atoi(string(digits))
	for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	var firstNonZeroIdx int
	for i := 0; i < len(digits); i++ {
		if digits[i] != '0' {
			firstNonZeroIdx = i
			break
		}
	}
	digits[firstNonZeroIdx], digits[0] = digits[0], digits[firstNonZeroIdx]
	lower, _ := strconv.Atoi(string(digits))

	// get all power of 2 numbers in this interval
	var candidates []int
	for i := 0; (1 << i) <= upper; i++ {
		a := 1 << i
		if a == lower || a == upper {
			return true
		}
		if a > lower {
			candidates = append(candidates, a)
		}
	}

	// compare n and candidates power of 2 numbers in a hash set manner
	mapEq := func(a, b map[byte]int) bool {
		if len(a) != len(b) {
			return false
		}

		for k, v := range a {
			if w, ok := b[k]; !ok || v != w {
				return false
			}
		}
		return true
	}

	digitsMap := map[byte]int{}
	for _, d := range digits {
		digitsMap[d]++
	}
	for _, c := range candidates {
		candidateMap, candidateDigits := map[byte]int{}, []byte(strconv.Itoa(c))
		for _, d := range candidateDigits {
			candidateMap[d]++
		}
		if mapEq(candidateMap, digitsMap) {
			return true
		}
	}
	return false
}
