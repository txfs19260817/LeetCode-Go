package _165_Compare_Version_Numbers

func compareVersion(version1 string, version2 string) int {
	for i, j := 0, 0; i < len(version1) || j < len(version2); {
		var x, y int
		for ; i < len(version1) && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++
		for ; j < len(version2) && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
	}
	return 0
}
