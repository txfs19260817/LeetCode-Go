package k

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPair(t *testing.T) {
	type args struct {
		songTimes [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example Case",
			args: args{
				songTimes: [][]string{
					{"Stairway to Heaven", "8:05"},
					{"Immigrant Song", "2:27"},
					{"Rock and Roll", "3:41"},
					{"Communication Breakdown", "2:29"},
					{"Good Times Bad Times", "2:48"},
					{"Hot Dog", "3:19"},
					{"The Crunge", "3:18"},
					{"Achilles Last Stand", "10:26"},
					{"Black Dog", "4:55"},
				},
			},
			want: []string{"Rock and Roll", "Hot Dog"},
		},
		{
			name: "No Pair",
			args: args{
				songTimes: [][]string{
					{"Song A", "1:00"},
					{"Song B", "2:00"},
					{"Song C", "3:00"},
				},
			},
			want: []string{},
		},
		{
			name: "Exact 7 Minutes Single Song (Should not match itself)",
			args: args{
				songTimes: [][]string{
					{"Long Song", "7:00"},
				},
			},
			want: []string{},
		},
		{
			name: "Pair with Zeros",
			args: args{
				songTimes: [][]string{
					{"Short", "0:00"}, // Assuming valid input format
					{"Long", "7:00"},
				},
			},
			want: []string{"Short", "Long"},
		},
		{
			name: "Multiple Pairs (Return First Found)",
			args: args{
				songTimes: [][]string{
					{"Pair1 A", "3:30"},
					{"Pair1 B", "3:30"},
					{"Pair2 A", "2:00"},
					{"Pair2 B", "5:00"},
				},
			},
			want: []string{"Pair1 A", "Pair1 B"},
		},
		{
			name: "Pair at Ends",
			args: args{
				songTimes: [][]string{
					{"Start", "1:00"},
					{"Middle", "4:00"},
					{"End", "6:00"},
				},
			},
			want: []string{"Start", "End"},
		},
		{
			name: "Invalid Format Handling",
			args: args{
				songTimes: [][]string{
					{"Bad Format", "invalid"},
					{"Another", "7:00"},
				},
			},
			want: []string{"Bad Format", "Another"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPair(tt.args.songTimes)
			// Sort or check existence because order might vary if map iteration was involved,
			// but here we iterate list so order is deterministic based on "seen" map.
			// My impl returns [first_seen, current].
			if !reflect.DeepEqual(got, tt.want) {
				// Try reverse order if order doesn't matter
				if len(got) == 2 && len(tt.want) == 2 {
					if got[0] == tt.want[1] && got[1] == tt.want[0] {
						return // Accept reverse order
					}
				}
				t.Errorf("findPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPlaylist(t *testing.T) {
	type args struct {
		initialSong string
		songs       []string
	}
	tests := []struct {
		name string
		args args
		want []string // One valid longest sequence
	}{
		{
			name: "Simple Chain",
			args: args{
				initialSong: "Hello World",
				songs: []string{
					"World of Warcraft",
					"Warcraft is Good",
					"Good Day",
				},
			},
			want: []string{"World of Warcraft", "Warcraft is Good", "Good Day"},
		},
		{
			name: "Branching (Choose Longest)",
			args: args{
				initialSong: "Start A",
				songs: []string{
					"A Path Long",
					"Long Way Home",
					"Home Sweet Home", // Path 1: Length 3
					"A Short",         // Path 2: Length 1
				},
			},
			want: []string{"A Path Long", "Long Way Home", "Home Sweet Home"},
		},
		{
			name: "No Match",
			args: args{
				initialSong: "Start End",
				songs: []string{
					"Start Again",
					"Hello World",
				},
			},
			want: []string{},
		},
		{
			name: "Cycle Handling (No Repeat)",
			args: args{
				initialSong: "A B",
				songs: []string{
					"B C",
					"C A",
					"A B", // Should not reuse if same song string, but wait, duplicate strings allowed in list?
					// "the song in the sequence should not repeat." usually means unique indices or unique content?
					// Given "initialSong is not given in the list", assuming list might have duplicates or not.
					// The problem says "the song in the sequence should not repeat."
					// If "A B" is in the list, it's a different instance than initial song.
					// But if "A B" appears twice in list, we can use both.
					// Assuming unique by index.
				},
			},
			// A B -> B C -> C A -> A B (from list). Length 3.
			want: []string{"B C", "C A", "A B"},
		},
		{
			name: "Complex Branching",
			args: args{
				initialSong: "Down By The River",
				songs: []string{
					"River of Dreams",
					"Dreams come true", // path 1 ends here (2)
					"River deep mountain high",
					"high hopes",
					"hopes and dreams", // path 2 ends here (3)
				},
			},
			want: []string{"River deep mountain high", "high hopes", "hopes and dreams"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPlaylist(tt.args.initialSong, tt.args.songs)
			if len(got) != len(tt.want) {
				t.Errorf("longestPlaylist() length = %v, want %v", len(got), len(tt.want))
			}
			// Check if `got` is a valid chain starting from initialSong
			if len(got) > 0 {
				// Verify chain validity
				// 1. Initial -> First
				_, initLast := getWords(tt.args.initialSong)
				firstFirst, _ := getWords(got[0])
				assert.Equal(t, initLast, firstFirst, "Chain broken: %s -> %s", tt.args.initialSong, got[0])
				// 2. i -> i+1
				for i := 0; i < len(got)-1; i++ {
					_, currLast := getWords(got[i])
					nextFirst, _ := getWords(got[i+1])
					assert.Equal(t, currLast, nextFirst, "Chain broken at index %d: %s -> %s", i, got[i], got[i+1])
				}
				// 3. Check if all songs are from the list (and unique indices used, implicit by algorithm)
				assert.Subset(t, tt.args.songs, got)
			}
		})
	}
}

// Helper for test verification
func getWords(s string) (string, string) {
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return "", ""
	}
	return parts[0], parts[len(parts)-1]
}
