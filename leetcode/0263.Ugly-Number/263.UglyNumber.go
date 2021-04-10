package _263_Ugly_Number

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		switch {
		case n%2 == 0:
			n /= 2
		case n%3 == 0:
			n /= 3
		case n%5 == 0:
			n /= 5
		default:
			return false
		}
	}
	return true
}
