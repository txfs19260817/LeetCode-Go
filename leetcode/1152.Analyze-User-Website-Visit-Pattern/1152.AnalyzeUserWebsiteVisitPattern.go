package leetcode

import "sort"

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	// sort 3 slices by timestamp
	type data struct {
		username  string
		timestamp int
		website   string
	}
	records := make([]data, len(username))
	for i := range username {
		records[i] = data{username[i], timestamp[i], website[i]}
	}
	sort.Slice(records, func(i, j int) bool { return records[i].timestamp < records[j].timestamp })

	// user to websites
	user2web := map[string][]string{}
	for _, a := range records {
		user2web[a.username] = append(user2web[a.username], a.website)
	}
	pattern2users, maxCnt, ans := map[[3]string][]string{}, 0, [3]string{}
	// extract patterns from each user's websites
	for user, sites := range user2web {
		// get all subarray with size 3
		for i := 0; i < len(sites)-2; i++ {
			for j := i + 1; j < len(sites)-1; j++ {
				for k := j + 1; k < len(sites); k++ {
					pattern := [3]string{sites[i], sites[j], sites[k]} // that subarray
					// if we have not recorded this pattern visited by this user, we should count it
					if len(pattern2users[pattern]) == 0 || pattern2users[pattern][len(pattern2users[pattern])-1] != user {
						pattern2users[pattern] = append(pattern2users[pattern], user) // record this user
						if c := len(pattern2users[pattern]); c > maxCnt {             // update max freqency
							maxCnt = c
							ans = pattern
						} else if c == maxCnt && (pattern[0] < ans[0] || pattern[0] == ans[0] && pattern[1] < ans[1] || pattern[0] == ans[0] && pattern[1] == ans[1] && pattern[2] < ans[2]) { // same freq but lexicographically smaller
							ans = pattern
						}
					}
				}
			}
		}
	}
	return ans[:] // array to slice
}
