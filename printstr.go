package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	strLen := len(s)

	for i := 0; i < strLen; i++ {
		z01.PrintRune(rune(s[i]))
	}
}
