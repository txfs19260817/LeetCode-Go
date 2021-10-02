package _405_Convert_a_Number_to_Hexadecimal

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || sb.Len() > 0 {
			var ch byte
			if val < 10 {
				ch = '0' + byte(val)
			} else {
				ch = 'a' + byte(val - 10)
			}
			sb.WriteByte(ch)
		}
	}
	return sb.String()
}

func toHex2(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}
	var ans []byte
	dict := map[int]byte{0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f'}
	for ; num != 0; num /= 16 {
		ans = append([]byte{dict[num%16]}, ans...)
	}
	return string(ans)
}
