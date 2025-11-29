package k

import (
	"math"
	"sort"
	"strconv"
)

/*
Resource Access Log

Question 1
Suppose we have an unsorted log file of accesses to web resources. Each log entry consists of an access
time, the ID of the user making the access, and the resource ID. The access time is represented as seconds
since 00:00:00, and all times are assumed to be in the same day. For example:

logs1 = [
["58523", "user_1", "resource_1"],
["62314", "user_2", "resource_2"],
["54001", "user_1", "resource_3"],
["200", "user_6", "resource_5"],
["215", "user_6", "resource_4"],
["54060", "user_2", "resource_3"],
["53760", "user_3", "resource_3"],
["58522", "user_22", "resource_1"],
["53651", "user_5", "resource_3"],
["2", "user_6", "resource_1"],
["100", "user_6", "resource_6"],
["400", "user_7", "resource_2"],
["100", "user_8", "resource_6"],
["54359", "user_1", "resource_3"],
]

Write a function that takes the logs and returns each users min and max access timestamp
*/

func getUserMinMax(logs [][]string) map[string][]int {
	userMinMax := make(map[string][]int)

	for _, log := range logs {
		timestamp, _ := strconv.Atoi(log[0])
		userID := log[1]

		if val, exists := userMinMax[userID]; exists {
			if timestamp < val[0] {
				val[0] = timestamp
			}
			if timestamp > val[1] {
				val[1] = timestamp
			}
		} else {
			userMinMax[userID] = []int{timestamp, timestamp}
		}
	}
	return userMinMax
}

/*
Question 2
Write a function that takes the logs and returns the resource with the highest number of accesses in any 5
minute window, together with how many accesses it saw.

Expected Output:
most_requested_resource(logs1) # => ('resource_3', 3)
Reason: resource_3 is accessed at 53760, 54001, and 54060
*/

func mostRequestedResource(logs [][]string) (string, int) {
	// Sort logs by timestamp first to process windows efficiently
	// Wait, logs are unsorted.
	// We need to group by resource, then sort accesses for each resource.

	resourceAccesses := make(map[string][]int)
	for _, log := range logs {
		timestamp, _ := strconv.Atoi(log[0])
		resourceID := log[2]
		resourceAccesses[resourceID] = append(resourceAccesses[resourceID], timestamp)
	}

	maxAccesses := 0
	topResource := ""

	for resourceID, timestamps := range resourceAccesses {
		sort.Ints(timestamps)

		// Sliding window of 5 minutes (300 seconds)
		for i := 0; i < len(timestamps); i++ {
			count := 0
			// Count how many subsequent accesses are within 300 seconds of timestamps[i]
			for j := i; j < len(timestamps); j++ {
				if timestamps[j]-timestamps[i] <= 300 {
					count++
				} else {
					break
				}
			}

			if count > maxAccesses {
				maxAccesses = count
				topResource = resourceID
			}
		}
	}

	return topResource, maxAccesses
}

/*
Question 3
Write a function that takes the logs as input, builds the transition graph and returns it as an adjacency list
with probabilities. Add START and END states.

Specifically, for each resource, we want to compute a list of every possible next step taken by any user,
together with the corresponding probabilities. The list of resources should include START but not END,
since by definition END is a terminal state.
*/

func transitionGraph(logs [][]string) map[string]map[string]float64 {
	// 1. Group logs by user and sort each user's logs by time
	type logEntry struct {
		timestamp int
		resource  string
	}
	userLogs := make(map[string][]logEntry)
	for _, log := range logs {
		ts, _ := strconv.Atoi(log[0])
		userLogs[log[1]] = append(userLogs[log[1]], logEntry{ts, log[2]})
	}

	for _, entries := range userLogs {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].timestamp < entries[j].timestamp
		})
	}

	// 2. Build transitions count
	transitions := make(map[string]map[string]int)
	totalTransitions := make(map[string]int)

	// Initialize START state
	transitions["START"] = make(map[string]int)

	for _, entries := range userLogs {
		// First resource is transition from START
		startRes := entries[0].resource
		transitions["START"][startRes]++
		totalTransitions["START"]++

		for i := 0; i < len(entries)-1; i++ {
			curr := entries[i].resource
			next := entries[i+1].resource

			if transitions[curr] == nil {
				transitions[curr] = make(map[string]int)
			}
			transitions[curr][next]++
			totalTransitions[curr]++
		}

		// Last resource transitions to END
		lastRes := entries[len(entries)-1].resource
		if transitions[lastRes] == nil {
			transitions[lastRes] = make(map[string]int)
		}
		transitions[lastRes]["END"]++
		totalTransitions[lastRes]++
	}

	// 3. Convert counts to probabilities
	result := make(map[string]map[string]float64)

	for state, nextStates := range transitions {
		result[state] = make(map[string]float64)
		total := float64(totalTransitions[state])

		for next, count := range nextStates {
			// Round to 3 decimal places to match example roughly (0.333)
			prob := float64(count) / total
			result[state][next] = math.Round(prob*1000) / 1000
		}
	}

	return result
}
