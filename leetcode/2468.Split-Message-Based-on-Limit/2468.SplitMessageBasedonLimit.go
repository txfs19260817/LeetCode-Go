package leetcode

import "fmt"

func splitMessage(message string, limit int) []string {
	for i, capacity, tailLen := 1, 0, 0; ; i++ {
		if i < 10 {
			tailLen = 5
		} else if i < 100 {
			if i == 10 {
				capacity -= 9
			}
			tailLen = 7
		} else if i < 1000 {
			if i == 100 {
				capacity -= 99
			}
			tailLen = 9
		} else {
			if i == 1000 {
				capacity -= 999
			}
			tailLen = 11
		}

		if tailLen >= limit {
			return nil
		}
		capacity += limit - tailLen
		if capacity < len(message) {
			continue
		}

		ans := make([]string, i)
		for j := range ans {
			tail := fmt.Sprintf("<%d/%d>", j+1, i)
			if j+1 == i {
				ans[j] = message + tail
			} else {
				m := limit - len(tail)
				ans[j] = message[:m] + tail
				message = message[m:]
			}
		}
		return ans
	}
}
