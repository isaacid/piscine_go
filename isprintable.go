package piscine

func IsPrintable(s string) bool {
	for _, c := range s {
		if !(rune(c) >= ' ' && rune(c) <= '~') {
			return false
		}
	}
	return true
}
