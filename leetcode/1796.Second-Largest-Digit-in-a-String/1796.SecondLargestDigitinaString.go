package _796_Second_Largest_Digit_in_a_String

func secondHighest(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if n := int(c - '0'); n > first {
				second = first
				first = n
			} else if n > second && n != first {
				second = n
			}
		}
	}
	return second
}
