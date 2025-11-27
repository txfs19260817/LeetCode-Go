package k

/*

We are working on a security system for a badged-access room in our company's building.

Given an ordered list of employees who used their badge to enter or exit the room, write a function that returns two collections:

All employees who didn't use their badge while exiting the room - they recorded an enter without a matching exit. (All employees are required to leave the room before the log ends.)

All employees who didn't use their badge while entering the room - they recorded an exit without a matching enter. (The room is empty when the log begins.)

Each collection should contain no duplicates, regardless of how many times a given employee matches the criteria for belonging to it.

records1 = [
["Paul", "enter"],
["Pauline", "exit"],
["Paul", "enter"],
["Paul", "exit"],
["Martha", "exit"],
["Joe", "enter"],
["Martha", "enter"],
["Steve", "enter"],
["Martha", "exit"],
["Jennifer", "enter"],
["Joe", "enter"],
["Curtis", "exit"],
["Curtis", "enter"],
["Joe", "exit"],
["Martha", "enter"],
["Martha", "exit"],
["Jennifer", "exit"],
["Joe", "enter"],
["Joe", "enter"],
["Martha", "exit"],
["Joe", "exit"],
["Joe", "exit"]
]
ENTER W/O EXIT EXIT W/O ENTER
Expected output: ["Steve", "Curtis", "Paul", "Joe"], ["Martha", "Pauline", "Curtis", "Joe"]

Other test cases:

records2 = [
["Paul", "enter"],
["Paul", "exit"],
]

Expected output: [], []

records3 = [
["Paul", "enter"],
["Paul", "enter"],
["Paul", "exit"],
["Paul", "exit"],
]

Expected output: ["Paul"], ["Paul"]

records4 = [
["Raj", "enter"],
["Paul", "enter"],
["Paul", "exit"],
["Paul", "exit"],
["Paul", "enter"],
["Raj", "enter"],
]

Expected output: ["Raj", "Paul"], ["Paul"]

All Test Cases:
mismatches(records1) => ["Steve", "Curtis", "Paul", "Joe"], ["Martha", "Pauline", "Curtis", "Joe"]
mismatches(records2) => [], []
mismatches(records3) => ["Paul"], ["Paul"]
mismatches(records4) => ["Raj", "Paul"], ["Paul"]

n: length of the badge records array

*/

func mismatches(records [][]string) (ans map[string][]string) {
	const ENTER, EXIT = "enter", "exit"
	inRoom := map[string]bool{}
	badEnter, badExit := map[string]bool{}, map[string]bool{}

	for _, record := range records {
		name, action := record[0], record[1]
		if action == ENTER {
			if inRoom[name] {
				badEnter[name] = true
			}
			inRoom[name] = true
		} else {
			if !inRoom[name] {
				badExit[name] = true
			}
			inRoom[name] = false
		}
	}

	for name, isIn := range inRoom {
		if isIn {
			badEnter[name] = true
		}
	}

	toList := func(m map[string]bool) []string {
		res := make([]string, 0, len(m))
		for k := range m {
			res = append(res, k)
		}
		return res
	}
	return map[string][]string{ENTER: toList(badEnter), EXIT: toList(badExit)}
}
