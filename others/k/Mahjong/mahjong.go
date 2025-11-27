package k

/*
You're creating a game with some amusing mini-games, and you've decided to make a simple variant of the game Mahjong.

In this variant, players have a number of tiles, each marked 0-9. The tiles can be grouped into pairs or triples of the same tile. For example, if a player has "333 444 66", the player's hand has a triple of 3s, a triple of 4s, and a pair of 6s. Similarly, "555 55 777" has a triple of 5s, a pair of 5s, and a triple of 7s.

A "complete hand" is defined as a collection of tiles where all the tiles can be grouped into any number of triples (zero or more) and exactly one pair, and each tile is used in exactly one triple or pair.

Write a function that takes a string representation of a collection of tiles in no particular order, and returns true or false depending on whether or not the collection represents a complete hand.

tiles1 = "11133555" # True. 111 33 555
tiles2 = "111333555" # False. There are three triples, 111 333 555 but no pair.
tiles3 = "00000111" # True. 000 00 111. Your pair and a triplet can be of the same value
# There is also no limit to how many of each tile there is.
tiles4 = "13233121" # True. Tiles are not guaranteed to be in order
tiles5 = "11223344555" # False. There cannot be more than one pair
tiles6 = "99999999" # True. You can have many of one tile
tiles7 = "999999999" # False.
tiles8 = "9" # False.
tiles9 = "99" # True. One pair.
tiles10 = "000022" # False.
tiles11 = "888889" # False. There cannot be any tiles left over.
tiles12 = "889" # False. There cannot be any tiles left over.
tiles13 = "88888844" # True. Two triples and one pair
tiles14 = "77777777777777" # True. Four triples and one pair
tiles15 = "1111111" # False.
tiles16 = "1111122222" # False.

Complexity Variable
N - Number of tiles in the input string
*/

func isCompleteHand(tiles string) bool {
	// Count frequencies of each tile
	counts := make(map[rune]int)
	for _, tile := range tiles {
		counts[tile]++
	}

	// A complete hand consists of any number of triples and exactly ONE pair.
	// This means exactly one tile type must have a count that leaves a remainder of 2 when divided by 3 (modulo 3 == 2).
	// All other tile types must have counts divisible by 3 (modulo 3 == 0).
	// Also, any tile type involved must have count >= 2 (for the pair) or >= 3 (for triples), but actually:
	// If count % 3 == 0, it's just triples. Valid counts: 3, 6, 9... (0 is not in map)
	// If count % 3 == 2, it's triples + 1 pair. Valid counts: 2, 5, 8...
	// If count % 3 == 1, it's impossible (e.g., 1, 4, 7...). 1 is obvious. 4 is 1 triple + 1 left over.

	pairCount := 0
	for _, count := range counts {
		if count%3 == 2 {
			// N Triple(s) + 1 pair
			pairCount++
		} else if count%3 == 1 {
			// e.g. 1, 4, 7 -> Impossible to form only triples and pairs
			return false
		}
	}

	// Must have exactly one pair component across all tile types
	return pairCount == 1
}

/*
You've decided to make a more advanced version of the Mahjong minigame, this time incorporating runs.

A run is a series of three consecutive tiles, like 123, 678, or 456. 0 counts as the lowest tile, so 012 is a valid run, but 890 is not.
A complete hand now consists of a pair, and any number (including zero) of triples and/or three-tile runs.
Write a function that returns whether or not a hand is complete under these new rules

 hand1 = "11123"        # True. 11 123
 hand2 = "12131"        # True. Also 11 123. Tiles are not necessarily sorted.
 hand3 = "11123455"     # True. 111 234 55
 hand4 = "11122334"     # True. 111 223 234
 hand5 = "11234"        # True. 11 234
 hand6 = "123456"       # False. Needs a pair
 hand7 = "111333555777" # True. 111 333 555 77
 hand8 = "11223344556677"  # True. 11 234 234 567 567 among others
 hand9 = "12233444556677"  # True. 123 234 44 567 567
 hand10 = "1123456789"  # False
 hand11 = "00123457"    # False
 hand12 = "0012345"     # False. A run is only three tiles
 hand13 = "11890"       # False. 890 is not a valid run
 hand14 = "99"          # True
 hand15 = "11122344"    # False
 All Test Cases
 advanced(hand1) => True
 advanced(hand2) => True
 advanced(hand3) => True
 advanced(hand4) => True
 advanced(hand5) => True
 advanced(hand6) => False
 advanced(hand7) => True
 advanced(hand8) => True
 advanced(hand9) => True
 advanced(hand10) => False
 advanced(hand11) => False
 advanced(hand12) => False
 advanced(hand13) => False
 advanced(hand14) => True
 advanced(hand15) => False
*/

func advanced(hand string) bool {
	counts := make(map[int]int)
	for _, tile := range hand {
		counts[int(tile-'0')]++
	}

	// Helper: Backtracking to see if tiles can be fully emptied by triples/runs
	var canSolve func() bool
	canSolve = func() bool {
		// 1. Find the smallest tile that exists
		start := -1
		for i := 0; i <= 9; i++ {
			if counts[i] > 0 {
				start = i
				break
			}
		}
		if start == -1 { // No tiles left -> Solved!
			return true
		}

		// 2. Try to form a Triple (e.g., 3,3,3)
		if counts[start] >= 3 {
			counts[start] -= 3
			if canSolve() {
				return true
			}
			counts[start] += 3 // Backtrack
		}

		// 3. Try to form a Run (e.g., 3,4,5)
		// 'start' must be the beginning of the run since it's the smallest available
		if start <= 7 && counts[start+1] > 0 && counts[start+2] > 0 {
			counts[start]--
			counts[start+1]--
			counts[start+2]--
			if canSolve() {
				return true
			}
			counts[start]++
			counts[start+1]++
			counts[start+2]++ // Backtrack
		}

		return false
	}

	// Try every possible tile as the "Pair"
	for i := 0; i <= 9; i++ {
		if counts[i] >= 2 {
			counts[i] -= 2
			if canSolve() {
				return true
			}
			counts[i] += 2 // Backtrack
		}
	}

	return false
}
