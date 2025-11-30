package k

import (
	"slices"
	"strings"
)

/*
You are reading a Build Your Own Story book. It is like a normal book except that choices on some pages affect the story, sending you to one of two options for your next page.
The choices are provided in a list, sorted by the page containing the choice, and each choice has two options of pages to go to next. In this example, you are on page 3, where there is a choice. Option 1 goes to page 14 and option 2 goes to page 2.
choices = [[3, 14, 2]] => [current_page, first_choice, second_choice]
You start reading at page 1 and read forward one page at a time unless you reach a choice or an ending.
You really love this book and so you decide to read all possible story sequences. You notice that you are reading some pages more than others, so you mostReadPage to find out which page you have read the most often when you read every storyline that leads to an ending.
You set some rules for your reading to avoid repeating pages too often. These rules are:
1) All storylines start at page 1.
2) Within any one storyline, you will never make the same choice twice (you may choose the other option)
3) If you reach a choice where you've already made both choices, you will not reach an ending, so this is not a valid storyline.
Given a list of endings and a list of choices with their destinations, return the page which was read the most often, as well as the number of times it was read. If multiple pages were read the same number of times, you may return any of them. If there are no valid storylines, return -1.
Example:
endings1 = [5, 10]
choices1_1 = [[3, 7, 9], [9, 10, 8]]
1 -> 2 -> 3(choice) -> 7 -> 8 -> 9(choice) -> 10(ending)
          |                      |
          |                      8 -> 9(choice, can't repeat 8) -> 10
          9(choice) -> 10(ending)
          |
          8 -> 9(choice, can't repeat 8) -> 10
All Storylines:
1->2->3->7->8->9->10
1->2->3->7->8->9->8->9->10
1->2->3->9->10
1->2->3->9->8->9->10
Page reads:
(1:4), (2:4), (3:4), (4:0), (5:0), (6:0), (7:2), (8:4), (9:6), (10:4)
Page with most reads = Page: 9, Reads: 6 (outputs can be in any format)
*/

func findMostReadPage(endings []int, choices [][]int) (int, int) {
	// ans
	mostReadPage, reads := -1, 0

	// convert input slices to sets
	endingSet := make(map[int]bool, len(endings))
	for _, e := range endings {
		endingSet[e] = true
	}
	choiceMap := make(map[int][2]int, len(choices))
	for _, c := range choices {
		choiceMap[c[0]] = [2]int{c[1], c[2]}
	}

	path, pageCount, choicesMade := []int{}, map[int]int{}, map[int]map[int]bool{}
	var dfs func(int)
	dfs = func(page int) {
		path = append(path, page)
		defer func() { path = path[:len(path)-1] }() // backtracking

		// Count all pages in this complete storyline if ending
		if endingSet[page] {
			for _, p := range path {
				pageCount[p]++
			}
			return
		}

		if twoChoices, ok := choiceMap[page]; ok {
			if choicesMade[page] == nil { // avoid nil map
				choicesMade[page] = make(map[int]bool, 2)
			}
			for i, nextPage := range twoChoices { // try each available option
				if !choicesMade[page][i] {
					choicesMade[page][i] = true
					dfs(nextPage)
					choicesMade[page][i] = false // backtracking
				}
			}
		} else {
			dfs(page + 1) // not a choice page, move onto the next page
		}
	}

	dfs(1)

	// Find the page with max reads
	for page, cnt := range pageCount {
		if reads < cnt {
			mostReadPage, reads = page, cnt
		}
	}
	return mostReadPage, reads
}

/*
Suppose you have a one-dimensional board of two colors of tiles. Red tiles can only move to the right, black tiles can only move to the left. A tile can move 1 space at a time. Either they move to an adjacent empty space, or they can jump over a single tile of the other color to an empty space.
Eg:
red = R
black = B
empty = _
R _ B _ => _ R B _ or
         R B _ _
R B _ _ => _ B R _
Given a start and end configuration represented as a list of strings, return a list of valid moves to get from start to end (doesn't need to be shortest), or None if none exist. Include the start and end states in the list of moves.
Example #1:
start_1 = ['R', '_', 'B', 'B']
end_1 = ['B', '_', 'B', 'R']
-> [
  ['R', '_', 'B', 'B']
  ['_', 'R', 'B', 'B']
  ['B', 'R', '_', 'B']
  ['B', 'R', 'B', '_']
  ['B', '_', 'B', 'R']
]
Example #2:
start_2 = ['R', 'R', '_']
end_2 = ['_', 'R', 'R']
-> [
  ['R', 'R', '_']
  ['R', '_', 'R']
  ['_', 'R', 'R']
]
All Test Cases:
validMoves(start_1, end_1)
validMoves(start_2, end_2)
n: number of spaces in the board
"""

Test cases:
start_1 = ["R","_","B","B"]
end_1 = ["B","_","B","R"]
start_2 = ["R", "R", "_"]
end_2 = ["_", "R", "R"]
*/

func validMoves(start, end []string) (res [][]string) {
	startStr, endStr := strings.Join(start, ""), strings.Join(end, "")
	if startStr == endStr {
		return [][]string{start}
	}

	parent := map[string]string{}
	reconstructPath := func() [][]string {
		var path []string
		for cur := endStr; len(cur) > 0; cur = parent[cur] {
			path = append(path, cur)
		}
		slices.Reverse(path)

		res = make([][]string, len(path))
		for i, p := range path {
			res[i] = make([]string, len(p))
			for j, c := range p {
				res[i][j] = string(c)
			}
		}
		return res
	}
	queue := []string{startStr}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, nextState := range getNextState(current) {
			if _, ok := parent[nextState]; !ok {
				parent[nextState] = current

				if nextState == endStr {
					return reconstructPath()
				}

				queue = append(queue, nextState)
			}
		}
	}
	return nil
}

func getNextState(state string) (res []string) {
	board := []byte(state)
	swap := func(i, j int) string {
		newBoard := slices.Clone(board)
		newBoard[i], newBoard[j] = newBoard[j], newBoard[i]
		return string(newBoard)
	}
	for i := 0; i < len(board); i++ {
		if board[i] == 'R' {
			if i+1 < len(board) && board[i+1] == '_' {
				res = append(res, swap(i, i+1))
			}
			if i+2 < len(board) && board[i+1] == 'B' && board[i+2] == '_' {
				res = append(res, swap(i, i+2))
			}
		} else if board[i] == 'B' {
			if i-1 >= 0 && board[i-1] == '_' {
				res = append(res, swap(i, i-1))
			}
			if i-2 >= 0 && board[i-1] == 'R' && board[i-2] == '_' {
				res = append(res, swap(i, i-2))
			}
		}
	}
	return
}

/*
You are going on a camping trip.
You and your friends are planning the hike to your campsite but there are various attractions you would like to see along the way. You are trying to plan a route where you can see all of your desired attractions without walking the same trail twice.
Trails are listed by the attractions they connect, all trails are 2-way trails and there can be multiple trails between 2 places.
Given a list of trails and a list of desired attractions, return whether there is a path which starts at the Parking Lot and ends at the Campsite which visits all of the desired attractions without using the same trail twice.
============================ ===================================
Examples:
===============================================================
Liberty Lake    Frozen Ocean   Eel Weir
    |            |        |    |      |
Parking Lot------Beaver Dam----Campsite
trails1 = [
    ["Beaver Dam",  "Frozen Ocean"], # First trail between Beaver Dam/Frozen Ocean
    ["Beaver Dam",  "Frozen Ocean"], # Second trail between Beaver Dam/Frozen Ocean
    ["Parking Lot", "Beaver Dam"],
    ["Parking Lot", "Liberty Lake"],
    ["Beaver Dam",  "Campsite"],
    ["Eel Weir",  "Campsite"],
    ["Eel Weir",  "Campsite"],
]
attractions1_1 = ["Frozen Ocean"] => True
Path: Parking Lot->Beaver Dam->Frozen Ocean->Beaver Dam->Campsite
attractions1_2 = ["Liberty Lake", "Beaver Dam"] => False
It is not possible to return from Liberty Lake so this path is not possible.
attractions1_3 = ["Eel Weir"] => True
Path: Parking Lot->Beaver Dam->Campsite->Eel Weir->Campsite
--------------------------------------------
                Liberty Lake
                |    |    |
Jeremy's Bay--Mason's Cabin--Parking Lot
      |                        |
    Horseshoe Falls----Mills Falls----Eel Weir--Outdoor Theater--Campsite
                                             \                   /
                                              --Hardwood Forest--
trails2 = [
    ["Mason's Cabin", "Liberty Lake"],
    ["Parking Lot", "Mill Falls"],
    ["Mason's Cabin", "Jeremy's Bay"],
    ["Eel Weir", "Hardwood Forest"],
    ["Outdoor Theater", "Campsite"],
    ["Jeremy's Bay", "Horseshoe Falls"],
    ["Mason's Cabin", "Parking Lot"],
    ["Mason's Cabin", "Liberty Lake"],
    ["Mill Falls", "Horseshoe Falls"],
    ["Mill Falls", "Eel Weir"],
    ["Hardwood Forest", "Campsite"],
    ["Eel Weir", "Outdoor Theater"],
    ["Liberty Lake", "Mason's Cabin"]
]
attractions2_1 = ["Jeremy's Bay", "Mason's Cabin", "Outdoor Theater"] #=> True
attractions2_2 = ["Outdoor Theater", "Eel Weir", "Hardwood Forest"] #=> False
attractions2_3 = ["Liberty Lake"] #=> True
attractions2_4 = ["Horseshoe Falls", "Eel Weir"] #=> True
All Test Cases:
sightseeing(trails1, attractions1_1) => True
sightseeing(trails1, attractions1_2) => False
sightseeing(trails1, attractions1_3) => True
sightseeing(trails2, attractions2_1) => True
sightseeing(trails2, attractions2_2) => False
sightseeing(trails2, attractions2_3) => True
sightseeing(trails2, attractions2_4) => True
Complexity Variable:
n = number of trails
*/

func sightseeing(trails [][]string, attractions []string) bool {
	const start, end string = "Parking Lot", "Campsite"
	type Edge struct {
		neighbor string
		trailIdx int
	}

	// build graph
	g := map[string][]Edge{}
	for i, e := range trails {
		a, b := e[0], e[1]
		g[a] = append(g[a], Edge{b, i})
		g[b] = append(g[b], Edge{a, i})
	}

	// attractions []string -> set
	attractionSet := map[string]bool{}
	for _, attraction := range attractions {
		attractionSet[attraction] = true
	}

	// Track visited attraction
	visitedAttr, usedTrails := map[string]bool{}, map[int]bool{}

	// Check if start is a attraction
	if attractionSet[start] {
		visitedAttr[start] = true
	}

	var dfs func(currentAttr string) bool
	dfs = func(currentAttr string) bool {
		// success cond: at end with all attraction visited
		if currentAttr == end && len(visitedAttr) == len(attractions) {
			return true
		}

		// try all trails from current location
		for _, edge := range g[currentAttr] {
			// skip visited trail
			if usedTrails[edge.trailIdx] {
				continue
			}
			usedTrails[edge.trailIdx] = true

			// Track if this is a new attraction visited
			if attractionSet[edge.neighbor] {
				visitedAttr[edge.neighbor] = true
			}

			if dfs(edge.neighbor) {
				return true
			}

			// Backtracking
			usedTrails[edge.trailIdx] = false
			if attractionSet[edge.neighbor] {
				delete(visitedAttr, edge.neighbor) // Use delete instead of assigning false because we check lengths above
			}
		}
		return false
	}

	return dfs(start)
}
