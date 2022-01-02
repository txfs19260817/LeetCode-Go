package _968_Number_of_Laser_Beams_in_a_Bank

func numberOfBeams(bank []string) int {
	var ans, prod int
	for _, s := range bank {
		var cnt int
		for _, n := range s {
			if n == '1' {
				cnt++
			}
		}
		if cnt == 0 {
			continue
		}
		ans += prod * cnt
		prod = cnt
	}
	return ans
}
