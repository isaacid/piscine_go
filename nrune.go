package piscine

func NRune(s string, n int) rune {
	n = n - 1
	runeSlice := []rune(s)

	lastIndex := len(runeSlice) - 1

	if n < 0 || n > lastIndex {
		return rune(0)
	}

	return runeSlice[n]
}
