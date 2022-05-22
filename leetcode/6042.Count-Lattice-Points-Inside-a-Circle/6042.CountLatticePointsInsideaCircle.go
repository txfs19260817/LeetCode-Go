package _042_Count_Lattice_Points_Inside_a_Circle

func countLatticePoints(circles [][]int) int {
	m := map[[2]int]bool{}
	for _, circle := range circles {
		o, r := [2]int{circle[0], circle[1]}, circle[2]
		m[o] = true
		for y := o[1]; y <= o[1]+r; y++ {
			for x := o[0]; x <= o[0]+r; x++ {
				if (x-o[0])*(x-o[0])+(y-o[1])*(y-o[1]) > r*r {
					continue
				}
				m[[2]int{x, y}] = true
				m[[2]int{x, o[1]*2 - y}] = true
				m[[2]int{o[0]*2 - x, y}] = true
				m[[2]int{o[0]*2 - x, o[1]*2 - y}] = true
			}
		}
	}
	return len(m)
}
