package _299_Bulls_and_Cows

import "strconv"

func getHint(secret string, guess string) string {
	bulls, cows, dictSecret, dictGuess := 0, 0, [10]int{}, [10]int{}
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			dictSecret[secret[i]-'0']++
			dictGuess[guess[i]-'0']++
		}
	}
	for i := 0; i <= 9; i++ {
		if dictSecret[i] < dictGuess[i] {
			cows += dictSecret[i]
		} else {
			cows += dictGuess[i]
		}
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
