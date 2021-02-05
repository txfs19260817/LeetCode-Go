package _844_Backspace_String_Compare

func backspaceCompare(S string, T string) bool {
	var sSkip, tSkip int
	for i, j := len(S)-1, len(T)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		for i >= 0 {
			if S[i] == '#' {
				sSkip++
				i--
			} else if sSkip > 0 {
				sSkip--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if T[j] == '#' {
				tSkip++
				j--
			} else if tSkip > 0 {
				tSkip--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
	}
	return true
}
