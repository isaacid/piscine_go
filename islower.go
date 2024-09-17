package piscine

func IsLower(s string) bool {
	for _, letter := range s {
		if !(rune(letter) >= 'a' && rune(letter) <= 'z') {
			return false
		}
	}
	return true
}
