package k

/*
Problem 1 – Customer Transaction Summary
Description

You are given a list of transactions, where each transaction is represented as an integer array of the form:

[customerId, eventId, amount]


Implement a function that, for a given customerId, returns:

The total number of distinct events this customer has, and

The total amount of money charged for this customer.

Your algorithm should run in linear time with respect to the number of transactions.

Function Signature (Go)
func SummarizeCustomer(transactions [][]int, customerID int) (eventCount int, totalAmount int)

Example 1

Input:

transactions = [
[1, 100, 300],
[2, 200, 300],
[1, 101, 300],
[1, 100, 300]
]
customerID = 1


Output:

eventCount = 2
totalAmount = 900


Explanation:

Customer 1 has events {100, 101} → 2 distinct events.

Total amount = 300 + 300 + 300 = 900.

Constraints

1 <= len(transactions) <= 10^5

len(transactions[i]) == 3

0 <= customerId, eventId, amount <= 10^9
*/

func SummarizeCustomer(transactions [][]int, customerID int) (eventCount int, totalAmount int) {
	idset := map[int]bool{}
	for _, transaction := range transactions {
		customerId, eventId, amount := transaction[0], transaction[1], transaction[2]
		if customerId == customerID {
			totalAmount += amount
			idset[eventId] = true
		}
	}
	return len(idset), totalAmount
}

/*
Problem 2 – Restaurant Recommendation From Friends
Description

You are given:

An integer array liked representing restaurant IDs that the user likes.

A 2D integer array friendsLiked, where friendsLiked[i] is the list of restaurant IDs liked by the i-th friend.

We want to recommend new restaurants to the user based on their friends’ preferences, using the following rules:

For each friend:

Count how many restaurants this friend likes in common with the user.

If the friend has at least 2 restaurants in common with the user, then:

Consider all restaurants that this friend likes but the user does not like yet.

For each such restaurant, increase its recommendation count by 1.

If the friend has fewer than 2 restaurants in common with the user, ignore this friend completely.

At the end, return:

The list of restaurant IDs with the highest recommendation count (i.e., liked by the largest number of such qualifying friends),

And that highest recommendation count.

Return the list of restaurant IDs sorted in ascending order.
If there is no restaurant to recommend, return an empty list and maxCount = 0.

Function Signature (Go)
func RecommendRestaurants(liked []int, friendsLiked [][]int) (recommended []int, maxCount int)

Example 1

Input:

liked = [1, 2, 3]

friendsLiked = [
[2, 3, 4],   // friend 0
[1, 4, 5],   // friend 1
[3, 6]       // friend 2
]


Output:

recommended = [4]
maxCount = 1


Explanation:

Friend 0: common with user = {2, 3} → 2 in common → qualifies.
New restaurants = {4} → restaurant 4 gets +1.

Friend 1: common with user = {1} → only 1 in common → ignored.

Friend 2: common with user = {3} → only 1 in common → ignored.

Only restaurant 4 is recommended by qualifying friends, with count 1.

Constraints

0 <= len(liked) <= 10^5

0 <= len(friendsLiked) <= 10^4

0 <= len(friendsLiked[i]) <= 10^4

Restaurant IDs are integers in range [0, 10^9].
*/

func RecommendRestaurants(liked []int, friendsLiked [][]int) (recommended []int, maxCount int) {
	likedSet := map[int]bool{}
	for _, l := range liked {
		likedSet[l] = true
	}

	res2cnt := map[int]int{}
	for _, friLikedIds := range friendsLiked {
		shared := 0
		for _, likedId := range friLikedIds {
			if likedSet[likedId] {
				shared++
			}
		}

		if shared >= 2 {
			for _, likedId := range friLikedIds {
				if !likedSet[likedId] {
					res2cnt[likedId]++
				}
			}
		}
	}

	for res, cnt := range res2cnt {
		if cnt > maxCount {
			maxCount = cnt
			recommended = []int{res}
		} else if cnt == maxCount {
			recommended = append(recommended, res)
		}
	}
	return
}
