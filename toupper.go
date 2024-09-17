package piscine

func isALetter1(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func ToUpper(s string) string {
	newStr := ""

	for _, letter := range s {
		if rune(letter) >= 'a' && rune(letter) <= 'z' {
			newStr += string(rune(letter) - 32)
		} else {
			newStr += string(letter)
		}
	}

	return newStr
}
