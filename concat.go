package piscine

func Concat(str1 string, str2 string) string {
	str1Len := len(str1)
	str2Len := len(str2)
	newStr := ""

	for i := 0; i < str1Len; i++ {
		newStr += string(str1[i])
	}

	for i := 0; i < str2Len; i++ {
		newStr += string(str2[i])
	}
	return newStr
}
