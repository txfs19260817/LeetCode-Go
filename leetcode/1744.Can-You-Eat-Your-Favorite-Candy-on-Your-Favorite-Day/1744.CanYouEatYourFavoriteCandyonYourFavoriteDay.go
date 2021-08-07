package _744_Can_You_Eat_Your_Favorite_Candy_on_Your_Favorite_Day

func canEat(candiesCount []int, queries [][]int) []bool {
	ans, smallerSum := make([]bool, len(queries)), make([]int, len(candiesCount))
	for i := 1; i < len(smallerSum); i++ {
		smallerSum[i] = candiesCount[i-1] + smallerSum[i-1]
	}
	for i, query := range queries {
		favoriteType, favoriteDay, dailyCap := query[0], query[1], query[2]
		earliest, latest := smallerSum[favoriteType]/dailyCap, smallerSum[favoriteType]+candiesCount[favoriteType]-1
		ans[i] = favoriteDay >= earliest && favoriteDay <= latest
	}
	return ans
}
