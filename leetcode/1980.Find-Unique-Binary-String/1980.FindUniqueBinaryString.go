package _980_Find_Unique_Binary_String

import "fmt"

func findDifferentBinaryString(nums []string) string {
	var ans string
	m := map[string]bool{}
	for _, num := range nums {
		m[num] = true
	}
	dec2bin := func(num, length int) string {
		bs, i := make([]byte, length), length-1
		for ; num > 0; num >>= 1 {
			bs[i] = byte('0' + num%2)
			i--
		}
		for ; i >= 0; i-- {
			bs[i] = '0'
		}
		return string(bs)
	}
	for i := 0; i < 1<<len(nums[0]); i++ {
		if s := dec2bin(i, len(nums[0])); !m[s] {
			ans = s
			break
		}
	}
	return ans
}

func findDifferentBinaryString2(nums []string) string {
	var ans string
	bin2dec := func(num string) (dec int) {
		for i, c := range num {
			if c == '1' {
				dec += 1 << (len(num) - 1 - i)
			}
		}
		return
	}
	m := map[int]bool{}
	for _, num := range nums {
		m[bin2dec(num)] = true
	}
	for i := 0; i < 1<<len(nums[0]); i++ {
		if !m[i] {
			ans = fmt.Sprintf("%b", i)
			for j := len(ans); j < len(nums[0]); j++ {
				ans = "0" + ans
			}
			break
		}
	}
	return ans
}
