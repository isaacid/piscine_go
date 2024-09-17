package piscine

func StrRev(s string) string {
	strLen := len(s)
	reversedStr := ""

	for i := strLen - 1; i >= 0; i-- {
		reversedStr += string(s[i])
	}

	return reversedStr
}
