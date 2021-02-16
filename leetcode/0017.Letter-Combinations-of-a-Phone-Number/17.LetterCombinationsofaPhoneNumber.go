package _017_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	keyboard := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ans = keyboard[rune(digits[0])]
	for _, d := range digits[1:] {
		line := keyboard[d]
		temp := make([]string, 0, len(ans)*len(line))
		for _, a := range ans {
			for _, b := range line {
				temp = append(temp, a+b)
			}
		}
		ans = temp
	}
	return ans
}
