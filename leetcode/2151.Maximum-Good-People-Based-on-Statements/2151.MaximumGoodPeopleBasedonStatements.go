package leetcode

func maximumGood(statements [][]int) int {
	var ans int
next:
	for bin := 0; bin < 1<<len(statements); bin++ { // bin is the binary representation of good/bad people possibility
		var goodCnt int // the number of good people
		for i, row := range statements {
			if bin>>i&1 == 1 { // if `i` is a good person
				for j, iStateToJ := range row { // verify i's statement
					if iStateToJ < 2 && bin>>j&1 != iStateToJ { // [i's statement regarding j] contradicts [j's identity] in the current situation
						continue next // start next bin
					}
				}
				goodCnt++
			}
		}
		if goodCnt > ans {
			ans = goodCnt
		}
	}
	return ans
}
