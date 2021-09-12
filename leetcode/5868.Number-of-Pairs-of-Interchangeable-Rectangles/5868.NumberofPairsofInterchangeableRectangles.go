package _868_Number_of_Pairs_of_Interchangeable_Rectangles

func interchangeableRectangles(rectangles [][]int) int64 {
	var ans int64
	m := map[float64]int64{}
	for _, rectangle := range rectangles {
		m[float64(rectangle[0])/float64(rectangle[1])]++
	}
	for _, n := range m {
		ans += n * (n - 1) / 2
	}
	return ans
}
