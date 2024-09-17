package piscine

func isNum(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundNumber := false

	for i, letter := range s {
		if letter == '-' && !foundNumber && (i == 0 || !isANum(rune(s[i-1]))) {
			sign = -1
		} else if isNum(letter) {
			foundNumber = true
			result = (result * 10) + int(letter-'0')
		}
	}

	return result * sign
}
