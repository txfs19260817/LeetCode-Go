package _178_Number_of_Valid_Words_for_Each_Puzzle

func findNumOfValidWords(words []string, puzzles []string) []int {
	str2uint := func(s string) (num uint) {
		for _, c := range s {
			num |= 1 << int(c-'a')
		}
		return
	}
	ans, wordBitmaps := make([]int, len(puzzles)), make([]uint, len(words))
	for i, word := range words {
		wordBitmaps[i] = str2uint(word)
	}
	for i, puzzle := range puzzles {
		pb := str2uint(puzzle)
		for _, wb := range wordBitmaps {
			if first := uint(1 << int(puzzle[0]-'a')); wb&first == first && pb&wb == wb { // a&b==b means a contains b in terms of bit 1
				ans[i]++
			}
		}
	}
	return ans
}

func findNumOfValidWords2(words []string, puzzles []string) []int {
	str2int := func(s string) (num int) {
		for _, c := range s {
			num |= 1 << int(c-'a')
		}
		return
	}
	ans, wordBit2cnt := make([]int, len(puzzles)), map[int]int{}
	for _, word := range words {
		wordBit2cnt[str2int(word)]++
	}
	for i, puzzle := range puzzles {
		pb := str2int(puzzle)
		for subsetEle := pb; subsetEle > 0; subsetEle = (subsetEle - 1) & pb {
			if subsetEle&(1<<int(puzzle[0]-'a')) != 0 {
				ans[i] += wordBit2cnt[subsetEle]
			}
		}
	}
	return ans
}
