package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if (rune(letter) >= 'a' && rune(letter) <= 'z') || (rune(letter) >= 'A' && rune(letter) <= 'Z') {
			count++
		}
	}
	return count
}
