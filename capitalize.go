package piscine

func isLowerCase(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

func isUpperCase(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func isAlphaNumeric(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isAlpha1(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	sLen := len(s)
	newStr := ""

	if isLowerCase(rune(s[0])) {
		newStr += string(rune(s[0]) - 32)
	} else {
		newStr += string(s[0])
	}

	for i := 1; i < sLen; i++ {
		if !(isAlphaNumeric(rune(s[i-1]))) && isLowerCase(rune(s[i])) {
			newStr += string(rune(s[i]) - 32)
		} else if isAlphaNumeric(rune(s[i-1])) && isUpperCase(rune(s[i])) {
			newStr += string(rune(s[i]) + 32)
		} else {
			newStr += string(s[i])
		}
	}

	return newStr
}
