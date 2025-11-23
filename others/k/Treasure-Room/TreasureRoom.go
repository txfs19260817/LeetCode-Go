package k

import "math"

/*
You are with your friends in a castle, where there are multiple rooms named after flowers.
Some of the rooms contain treasures - we call them the treasure rooms.

Each room contains a single instruction that tells you which room to go to next.

*** instructions_1 and treasure_rooms_1 ***

lily* ---------      daisy  sunflower
               |        |     |
               v        v     v
jasmin --> tulip*      violet* ----> rose* -->
            ^    |      ^             ^       |
            |    |      |             |       |
            ------    iris            ---------

* denotes a treasure room, e.g., rose is a treasure room, but jasmin isn't.

This is given as a list of pairs of (source_room, destination_room)

instructions_1 = [
    ["jasmin", "tulip"],
    ["lily", "tulip"],
    ["tulip", "tulip"],
    ["rose", "rose"],
    ["violet", "rose"],
    ["sunflower", "violet"],
    ["daisy", "violet"],
    ["iris", "violet"]
]

treasure_rooms_1 = ["lily", "tulip", "violet", "rose"]

Write a function that takes two parameters as input:
* a list of instructions represented as pairs of (source_room, destination_room), and
* a list containing the treasure rooms,

and returns a collection of all the rooms that satisfy the following two conditions:
* at least two *other* rooms have instructions pointing to this room
* this room's instruction immediately points to a treasure room

filter_rooms(instructions_1, treasure_rooms_1) => ["tulip", "violet"]
* tulip can be accessed from rooms lily and jasmin. Tulip's instruction points to a treasure room (tulip itself)
* violet can be accessed from daisy, sunflower and iris. Violet's instruction points to a treasure room (rose)

Additional inputs

treasure_rooms_2 = ["lily", "jasmin", "violet"]

filter_rooms(instructions_1, treasure_rooms_2) => []
* none of the rooms reachable from tulip or violet are treasure rooms

*** instructions_2 and treasure_rooms_3 ***

lily ---------          --------
              |          |      |
              v          v      |
jasmin --> tulip ---> violet*--^

instructions_2 = [
    ["jasmin", "tulip"],
    ["lily", "tulip"],
    ["tulip", "violet"],
    ["violet", "violet"]
]

treasure_rooms_3 = ["violet"]

filter_rooms(instructions_2, treasure_rooms_3) => [tulip]
* tulip can be accessed from rooms lily and jasmin. Tulip's instruction points to a treasure room (violet)

All the test cases:
filter_rooms(instructions_1, treasure_rooms_1)    => ["tulip", "violet"]
filter_rooms(instructions_1, treasure_rooms_2)    => []
filter_rooms(instructions_2, treasure_rooms_3)    => [tulip]

Complexity Analysis variables:
T: number of treasure rooms
I: number of instructions given
*/

func filterRooms(instructions [][]string, treasureRooms []string) (ans []string) {
	treasureRoomsSet := map[string]bool{}
	for _, room := range treasureRooms {
		treasureRoomsSet[room] = true
	}

	inDegree, nextRoom := map[string]int{}, map[string]string{}
	for _, instruction := range instructions {
		from, to := instruction[0], instruction[1]
		nextRoom[from] = to
		if from != to {
			inDegree[to]++
		}
	}

	for room, inDeg := range inDegree {
		if inDeg > 1 && len(nextRoom[room]) > 0 && treasureRoomsSet[nextRoom[room]] {
			ans = append(ans, room)
		}
	}
	return
}

/*
We are playing a game where the player needs to follow instructions to find a treasure.

There are multiple rooms, aligned in a straight line, labeled sequentially from 0. Each room contains one
instruction, given as a positive integer.

An instruction directs the player to move forward a specific number of rooms. The last instruction is "9" by
convention, and can be ignored (there's no room to move after the last room).

The player starts the game in room number 0 and has to reach the treasure which is in the last room. The
player is given an amount of money to start the game with. She must use this money wisely to get to the
treasure as fast as possible.

The player can follow the instruction or pay 1, the instruction "2" may be changed to "1" or "3". A player
cannot pay more than $1 to change the value of an instruction by more than one unit.

Write a function that takes a list of instructions and a total amount of money as input and returns the
minimum number of instructions needed to reach the treasure room, or None/null/-1 if the treasure cannot
be reached.

Examples

	Note: The updated instructions are marked with *.

Example 1
instructions_2_1 = [1, 1, 1, 9]
With $0, the player would follow 3 instructions: Instructions: [ 1, 1, 1, 9]
Itinerary: [ 1, 1, 1, 9] ^ ^ ^ ^
With $1, the player would reach the treasure in 2 instructions: she could change, for example, the first instruction to 2.
Instructions: [ 1, 1, 1, 9]
Itinerary: [ *2, 1, 1, 9] ^ ^ ^
Example 2
instructions_2_2 = [1, 1, 2, 9]
With $0 as the initial amount, the treasure is not reachable.
With $1 (or more) as the initial amount, the treasure can be reached in 2 instructions.
Instructions: [ 1, 1, 2, 9] Itinerary: [ 1, *2, 2, 9]
*/
func minInstructionsToTreasure(instructions []int, money int) int {
	ans := math.MaxInt
	bestSteps := make([][]int, len(instructions))
	for i := range bestSteps {
		bestSteps[i] = make([]int, money+1)
		for j := range bestSteps[i] {
			bestSteps[i][j] = math.MaxInt
		}
	}
	var dfs func(int, int, int)
	dfs = func(index int, remaining int, steps int) {
		if index < 0 || index >= len(instructions) || remaining < 0 || steps >= ans {
			return
		}
		if steps >= bestSteps[index][remaining] {
			return
		}
		bestSteps[index][remaining] = steps
		if instructions[index] == 9 || index == len(instructions)-1 {
			ans = min(ans, steps)
			return
		}
		dfs(index+instructions[index], remaining, steps+1)
		dfs(index+instructions[index]+1, remaining-1, steps+1)
		if instructions[index] > 1 {
			dfs(index+instructions[index]-1, remaining-1, steps+1)
		}
	}
	dfs(0, money, 0)
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}
