package _957_Adding_Spaces_to_a_String

func addSpaces(s string, spaces []int) string {
	ans, offset := make([]rune, len(s)+len(spaces)), 0
	for i, space := range spaces {
		ans[space+i] = ' '
	}
	for _, c := range s {
		for ans[offset] != 0 {
			offset++
		}
		ans[offset] = c
	}
	return string(ans)
}
