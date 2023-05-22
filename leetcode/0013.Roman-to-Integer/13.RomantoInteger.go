package leetcode

func romanToInt(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		var singleFlag bool
		if i+1 < len(s) {
			switch s[i : i+2] {
			case "IV":
				ans += 4
				i++
			case "IX":
				ans += 9
				i++
			case "XL":
				ans += 40
				i++
			case "XC":
				ans += 90
				i++
			case "CD":
				ans += 400
				i++
			case "CM":
				ans += 900
				i++
			default:
				singleFlag = true
			}
		} else {
			singleFlag = true
		}
		if singleFlag {
			switch s[i] {
			case 'I':
				ans += 1
			case 'V':
				ans += 5
			case 'X':
				ans += 10
			case 'L':
				ans += 50
			case 'C':
				ans += 100
			case 'D':
				ans += 500
			case 'M':
				ans += 1000
			}
		}
	}
	return ans
}
