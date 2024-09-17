package piscine

func isALetter(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func isANum(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !((isALetter(rune(letter))) || (isANum(rune(letter)))) {
			return false
		}
	}

	return true
}
