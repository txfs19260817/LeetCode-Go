package k

import (
	"slices"
	"strconv"
	"strings"
)

/*
We have a catalog of song titles (and their lengths) that we play at a local radio station.  We have been asked to play two of those songs in a row, and they must add up to exactly seven minutes long.


Given a list of songs and their durations, write a function that returns the names of any two distinct songs that add up to exactly seven minutes.  If there is no such pair, return an empty collection.

song_times_1 = [
    ("Stairway to Heaven", "8:05"), ("Immigrant Song", "2:27"),
    ("Rock and Roll", "3:41"), ("Communication Breakdown", "2:29"),
    ("Good Times Bad Times", "2:48"), ("Hot Dog", "3:19"),
    ("The Crunge", "3:18"), ("Achilles Last Stand", "10:26"),
    ("Black Dog", "4:55")
]
find_pair(song_times_1) => ["Rock and Roll", "Hot Dog"] (3:41 + 3:19 = 7:00)
*/

func findPair(songTimes [][]string) []string {
	const targetSec int = 7 * 60
	timeToSec := func(t string) int {
		parts := strings.Split(t, ":")
		if len(parts) != 2 {
			return 0
		}
		min, _ := strconv.Atoi(parts[0])
		sec, _ := strconv.Atoi(parts[1])
		return min*60 + sec
	}
	// Map stores duration in seconds -> song name
	seen := map[int]string{}

	for _, songTime := range songTimes {
		if len(songTime) < 2 {
			continue
		}
		songName, timeStr := songTime[0], songTime[1]
		curSec := timeToSec(timeStr)
		neededSec := targetSec - curSec

		if otherName, ok := seen[neededSec]; ok {
			// Found a pair! Return them. Order doesn't strictly matter based on prompt,
			// but returning them as found or in specific order is fine.
			// The example output is ["Rock and Roll", "Hot Dog"], where "Rock and Roll" came first in the list.
			// So let's return [otherName, songName] to match the "first found, second found" order implicitly.
			return []string{otherName, songName}
		}
		// Store current song for future matches
		seen[curSec] = songName
	}
	return []string{}
}

/*
Given list of songs as String[] and initialSong, find the longest sequence of songs as List<String> with the following consideration.
- initialSong is not given in the the list of songs.
- the next song in sequence should start with the same word as last word in the previous song.
- if there are multiple such songs, select any one
- the song in the sequence should not repeat.
- words in song is separated by space.
*/
func longestPlaylist(initialSong string, songs []string) []string {
	n := len(songs)
	firstWords, lastWords := make([]string, n), make([]string, n)
	for i, song := range songs {
		words := strings.Fields(song)
		firstWords[i], lastWords[i] = words[0], words[len(words)-1]
	}

	firstWord2Indices := map[string][]int{}
	for i, word := range firstWords {
		firstWord2Indices[word] = append(firstWord2Indices[word], i)
	}

	visitedIdx := make([]bool, n)
	var bestPath []int
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		visitedIdx[idx] = true
		defer func() { visitedIdx[idx] = false }() // Backtrack; path slice is discarded on return (passed by value)

		path = append(path, idx)
		// Update bestPath if this path is longer.
		if len(bestPath) < len(path) {
			bestPath = slices.Clone(path)
		}

		// Continue to any song whose first word == last word of current song.
		nextWord := lastWords[idx]
		for _, j := range firstWord2Indices[nextWord] {
			if !visitedIdx[j] {
				dfs(j, path)
			}
		}
	}

	// Start DFS from all songs that start with the last word of initialSong.
	initialSongParts := strings.Fields(initialSong)
	beginWord := initialSongParts[len(initialSongParts)-1]
	for _, idx := range firstWord2Indices[beginWord] {
		if !visitedIdx[idx] {
			dfs(idx, nil)
		}
	}

	// Convert bestPath indices to []string.
	result := make([]string, len(bestPath))
	for i, idx := range bestPath {
		result[i] = songs[idx]
	}
	return result
}
