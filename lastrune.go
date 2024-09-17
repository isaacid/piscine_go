package piscine

func LastRune(s string) rune {
	runeSlice := []rune(s)

	lastIndex := len(runeSlice) - 1

	return runeSlice[lastIndex]
}
