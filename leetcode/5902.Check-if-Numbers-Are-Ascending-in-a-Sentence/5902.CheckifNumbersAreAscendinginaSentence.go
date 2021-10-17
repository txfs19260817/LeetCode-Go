package _902_Check_if_Numbers_Are_Ascending_in_a_Sentence

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	parts, num := strings.Split(s, " "), -1
	for _, part := range parts {
		a, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		if a <= num {
			return false
		}
		num = a
	}
	return true
}
