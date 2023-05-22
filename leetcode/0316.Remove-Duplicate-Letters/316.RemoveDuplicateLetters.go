package leetcode

import "sort"

func removeDuplicateLetters(s string) string {
	var stack []rune
	var ch2indices [26][]int
	for i, r := range s {
		ch2indices[int(r-'a')] = append(ch2indices[int(r-'a')], i)
	}
	hashSet := map[rune]bool{}
	for i, r := range s {
		if hashSet[r] {
			continue
		}
		/*
			The conditions for popping from mono-stack are:
			  - The character is greater than the current characters
			  - The character can be removed because it occurs later on
		*/
		for len(stack) > 0 && stack[len(stack)-1] > r && sort.SearchInts(ch2indices[int(stack[len(stack)-1]-'a')], i) < len(ch2indices[int(stack[len(stack)-1]-'a')]) {
			delete(hashSet, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, r)
		hashSet[r] = true
	}
	return string(stack)
}
