package piscine

func ConcatParams(args []string) string {
	argsLen := len(args)
	newStr := ""

	for i := 0; i < argsLen; i++ {
		for _, letter := range args[i] {
			newStr += string(letter)
		}
		if !(i == argsLen-1) {
			newStr += string('\n')
		}
	}

	return newStr
}
