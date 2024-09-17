package piscine

func CompareA(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	minIndex := lenA

	if lenA > lenB {
		minIndex = lenB
	}

	for i := 0; i < minIndex; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	}

	return 0
}

func Index(s string, toFind string) int {
	sLen := len(s)
	toFindLen := len(toFind)

	if toFindLen < 0 || toFindLen > sLen {
		return -1
	}

	for i := 0; i < sLen-toFindLen; i++ {
		if CompareA(s[i:i+toFindLen], toFind) == 0 {
			return i
		}
	}

	return -1
}
