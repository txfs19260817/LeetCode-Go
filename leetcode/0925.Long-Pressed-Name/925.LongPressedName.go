package _925_Long_Pressed_Name

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	var i, j int
	for n, m := 0, 0; i < len(name) && j < len(typed); n, m = 0, 0 {
		if name[i] != typed[j] {
			return false
		}
		for i+1 < len(name) && name[i+1] == name[i] {
			i++
			n++
		}
		i++
		n++
		for j+1 < len(typed) && typed[j+1] == typed[j] {
			j++
			m++
		}
		j++
		m++
		if m < n {
			return false
		}
	}
	if i < len(name) || j < len(typed) {
		return false
	}
	return true
}
