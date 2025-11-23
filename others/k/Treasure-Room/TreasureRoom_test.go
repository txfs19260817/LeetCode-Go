package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterRooms(t *testing.T) {
	instructions1 := [][]string{
		{"jasmin", "tulip"},
		{"lily", "tulip"},
		{"tulip", "tulip"},
		{"rose", "rose"},
		{"violet", "rose"},
		{"sunflower", "violet"},
		{"daisy", "violet"},
		{"iris", "violet"},
	}
	treasure1 := []string{"lily", "tulip", "violet", "rose"}

	treasure2 := []string{"lily", "jasmin", "violet"}

	instructions2 := [][]string{
		{"jasmin", "tulip"},
		{"lily", "tulip"},
		{"tulip", "violet"},
		{"violet", "violet"},
	}
	treasure3 := []string{"violet"}

	// New inputs

	// Empty instructions: nothing should qualify.
	instructionsEmpty := [][]string{}
	treasureEmpty := []string{}

	// Self-loop plus one external predecessor on the same node:
	// "b" has predecessors {a, b} → only 1 *other* room → should NOT qualify.
	instructions3 := [][]string{
		{"a", "b"},
		{"b", "b"},
	}
	treasure4 := []string{"b"}

	// Node "b" has 2 predecessors (a, c), but its instruction points to "c",
	// and the treasure is "d" reachable only via c → non-immediate treasure,
	// so "b" must NOT qualify.
	instructions4 := [][]string{
		{"a", "b"},
		{"c", "b"},
		{"b", "c"},
		{"c", "d"},
		{"d", "d"},
	}
	treasure5 := []string{"d"}

	// Node "b" has three *other* predecessors (a, c, d) and a self-loop.
	// "b" is a treasure and its instruction points to itself (a treasure),
	// so "b" SHOULD qualify.
	instructions5 := [][]string{
		{"a", "b"},
		{"c", "b"},
		{"d", "b"},
		{"b", "b"},
	}
	treasure6 := []string{"b"}

	// Destination treasure "t" doesn’t need to have an outgoing instruction
	// that satisfies any condition. Only "y" has >= 2 other predecessors and
	// points directly to treasure "t".
	instructions6 := [][]string{
		{"x", "y"},
		{"y", "t"},
		{"z", "y"},
		{"w", "y"},
		{"t", "t"},
	}
	treasure7 := []string{"t"}

	tests := []struct {
		name         string
		instructions [][]string
		treasures    []string
		expected     []string
	}{
		// Original cases
		{"case1_baseExample", instructions1, treasure1, []string{"tulip", "violet"}},
		{"case2_noImmediateTreasureFromPopularRooms", instructions1, treasure2, []string{}},
		{"case3_pointsToTreasureOtherRoom", instructions2, treasure3, []string{"tulip"}},

		// New cases
		{"case4_emptyInstructions", instructionsEmpty, treasure1, []string{}},
		{"case5_noTreasureRooms", instructions1, treasureEmpty, []string{}},
		{"case6_selfLoopNotCountedAsOtherPredecessor", instructions3, treasure4, []string{}},
		{"case7_nonImmediateTreasureShouldFail", instructions4, treasure5, []string{}},
		{"case8_multiplePredecessorsAndSelfLoop", instructions5, treasure6, []string{"b"}},
		{"case9_destinationDoesNotNeedOwnInstruction", instructions6, treasure7, []string{"y"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterRooms(tt.instructions, tt.treasures)
			assert.ElementsMatch(t, tt.expected, got)
		})
	}
}

func Test_minInstructionsToTreasure(t *testing.T) {
	type args struct {
		instructions []int
		money        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Provided examples
		{"example1_noMoney", args{[]int{1, 1, 1, 9}, 0}, 3},
		{"example1_money1", args{[]int{1, 1, 1, 9}, 1}, 2},

		{"example2_noMoney_unreachable", args{[]int{1, 1, 2, 9}, 0}, -1},
		{"example2_money1", args{[]int{1, 1, 2, 9}, 1}, 2},

		// --- New cases ---

		// Already at treasure: length 1, starting room is the last room.
		{"alreadyAtTreasure_len1", args{[]int{9}, 0}, 0},

		// Two rooms: basic jump, no need for money.
		{"twoRooms_basic", args{[]int{1, 9}, 0}, 1},

		// Simple path where no changes are needed.
		// 0 -> 2 -> 3
		{"simplePath_noChangeNeeded", args{[]int{2, 2, 1, 9}, 0}, 2},

		// Same instructions, but with money you can jump directly:
		// change first 2 -> 3, 0 -> 3
		{"simplePath_changeFirst_jump", args{[]int{2, 2, 1, 9}, 1}, 1},

		// Must change at some point to be reachable; without money it's impossible.
		// 0 -> 2 (then 2 -> 4 out of bounds)
		{"changeFirst_reachInOneStep_unreachableWithoutMoney", args{[]int{2, 2, 2, 9}, 0}, -1},

		// With 1 money: change first 2 -> 3, 0 -> 3 (1 step).
		{"changeFirst_reachInOneStep_money1", args{[]int{2, 2, 2, 9}, 1}, 1},

		// A case where money gives a shorter path, but without money it’s still reachable.
		// No money: 0 -> 1 -> 3 -> 4 (3 steps).
		// With 1 money: change first 1 -> 2, 0 -> 2 -> 4 (2 steps).
		{"shorterPathWithoutMoney", args{[]int{1, 2, 2, 1, 9}, 0}, 3},
		{"shorterPathWithMoney", args{[]int{1, 2, 2, 1, 9}, 1}, 2},

		// Requires *two* paid changes for the optimal path.
		// instructions = [1,1,3,1,9]
		// money=0: unreachable.
		// money=1: 0->1->2->4 (change at 2: 3 -> 2) => 3 steps.
		// money=2: change at 0 (1->2): 0->2, change at 2 (3->2): 2->4 => 2 steps.
		{"requiresTwoPayments_unreachableWith0", args{[]int{1, 1, 3, 1, 9}, 0}, -1},
		{"requiresTwoPayments_money1", args{[]int{1, 1, 3, 1, 9}, 1}, 3},
		{"requiresTwoPayments_money2", args{[]int{1, 1, 3, 1, 9}, 2}, 2},

		// More money than necessary should not change the minimum number of steps.
		// Same as example1, but with extra money.
		{"example1_moreMoneySameResult", args{[]int{1, 1, 1, 9}, 3}, 2},

		// Unreachable even with lots of money because every possible jump
		// (5, 4, or 6) goes beyond the last room from index 0.
		{"unreachableEvenWithLotsOfMoney", args{[]int{5, 5, 5, 9}, 10}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				minInstructionsToTreasure(tt.args.instructions, tt.args.money),
				"minInstructionsToTreasure(%v, %v)",
				tt.args.instructions,
				tt.args.money,
			)
		})
	}
}

func Test_minInstructionsToTreasure_Huge(t *testing.T) {
	const n = 100_000
	instructions := make([]int, n)
	for i := 0; i < n-1; i++ {
		instructions[i] = 1
	}
	instructions[n-1] = 9

	// No money: must walk one by one
	got := minInstructionsToTreasure(instructions, 0)
	want := n - 1
	assert.Equal(t, want, got)

	// With money: you can try to see if your implementation still finishes fast.
	// (Exact "want" depends on your strategy; this is mostly a perf test.)
	got2 := minInstructionsToTreasure(instructions, 50_000)
	t.Logf("with money=50_000, steps=%d", got2)
}
