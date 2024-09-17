package piscine

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(rune(c) >= '0' && rune(c) <= '9') {
			return false
		}
	}
	return true
}
