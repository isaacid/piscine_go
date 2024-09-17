package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
func main() {

	argsArr := os.Args
	argsArrLen := len(argsArr)

	if argsArrLen != 4 {
		return
	}

	str := argsArr[1]
	char := argsArr[2]
	char_replace := argsArr[3]
	newStr := ""

	for _, letter := range str {
		if string(letter) == char {
			newStr += char_replace
		} else {
			newStr += string(letter)
		}
	}

	PrintStr(newStr)
}
