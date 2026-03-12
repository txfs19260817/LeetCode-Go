package leetcode

func russianPeasantMultiply(a, b int) int {
	return russianPeasantMultiplyIterative(a, b)
}

func russianPeasantMultiplyIterative(a, b int) int {
	a, b, negative := normalizeRussianPeasantSigns(a, b)

	result := 0
	for a > 0 {
		if a&1 == 1 {
			result += b
		}
		a >>= 1
		b <<= 1
	}

	if negative {
		return -result
	}
	return result
}

func russianPeasantMultiplyRecursive(a, b int) int {
	a, b, negative := normalizeRussianPeasantSigns(a, b)
	result := russianPeasantMultiplyRecursivePositive(a, b)
	if negative {
		return -result
	}
	return result
}

func russianPeasantMultiplyRecursivePositive(a, b int) int {
	if a == 0 {
		return 0
	}
	if a&1 == 1 {
		return b + russianPeasantMultiplyRecursivePositive(a>>1, b<<1)
	}
	return russianPeasantMultiplyRecursivePositive(a>>1, b<<1)
}

func normalizeRussianPeasantSigns(a, b int) (int, int, bool) {
	negative := false
	if a < 0 {
		a = -a
		negative = !negative
	}
	if b < 0 {
		b = -b
		negative = !negative
	}
	return a, b, negative
}
