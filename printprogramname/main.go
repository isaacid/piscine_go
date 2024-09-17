package main

import (
	"os"

	"github.com/01-edu/z01"
)

func findlastSlash(s string) int {
	lastIndex := 0

	for i, c := range s {
		if c == '/' {
			lastIndex = i
		}
	}
	return lastIndex
}

func main() {
	arg := os.Args[0]
	startingIndex := findlastSlash(arg) + 1 // start after the /

	for _, c := range arg[startingIndex:] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
