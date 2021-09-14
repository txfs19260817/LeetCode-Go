package _447_Number_of_Boomerangs

func numberOfBoomerangs(points [][]int) int {
	var ans int
	getDist := func(i, j int) int {
		x, y := points[i][0]-points[j][0], points[i][1]-points[j][1]
		return x*x + y*y
	}
	for i := range points {
		cnt := map[int]int{}
		for j := range points {
			cnt[getDist(i, j)]++
		}
		for _, n := range cnt {
			ans += n * (n - 1)
		}
	}
	return ans
}
