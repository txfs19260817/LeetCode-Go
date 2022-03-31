package _028_Count_Collisions_on_a_Road

func countCollisions(directions string) int {
	var ans, carsFromR, i int
	for i < len(directions) && directions[i] == 'L' {
		i++
	}
	for ; i < len(directions); i++ {
		if directions[i] == 'R' {
			carsFromR++
			continue
		}
		ans += carsFromR
		if directions[i] == 'L' {
			ans++
		}
		carsFromR = 0
	}
	return ans
}
