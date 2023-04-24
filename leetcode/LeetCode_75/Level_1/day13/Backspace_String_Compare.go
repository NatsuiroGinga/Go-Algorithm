package day13

func backspaceCompare(s string, t string) bool {
	i, j, back := len(s)-1, len(t)-1, 0

	for {
		back = 0
		for i >= 0 && (back > 0 || s[i] == '#') {
			if s[i] == '#' {
				back++
			} else {
				back--
			}
			i--
		}

		back = 0
		for j >= 0 && (back > 0 || t[i] == '#') {
			if t[i] == '#' {
				back++
			} else {
				back--
			}
			j--
		}

		if i >= 0 && j >= 0 && s[i] == t[i] {
			j--
			i--
		} else {
			break
		}
	}

	return i == -1 && j == -1
}
