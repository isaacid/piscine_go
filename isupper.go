package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if !(rune(letter) >= 'A' && rune(letter) <= 'Z') {
			return false
		}
	}
	return true
}
