package leetcode

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	m := map[int]int{}
	for _, barcode := range barcodes {
		m[barcode]++
	}

	Bs := make([]B, 0, len(m))
	for k, v := range m {
		Bs = append(Bs, B{k, v})
	}

	sort.Slice(Bs, func(i, j int) bool {
		return Bs[i].count > Bs[j].count
	})

	i := -2
	for _, b := range Bs {
		for j := 0; j < b.count; j++ {
			i += 2
			if i >= len(barcodes) {
				i = 1
			}
			barcodes[i] = b.barcode
		}
	}
	return barcodes
}

type B struct {
	barcode int
	count   int
}
