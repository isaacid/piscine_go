package piscine

func ToLower(s string) string {
	newStr := ""

	for _, letter := range s {
		if rune(letter) >= 'A' && rune(letter) <= 'Z' {
			newStr += string(rune(letter) + 32)
		} else {
			newStr += string(letter)
		}
	}

	return newStr
}
