package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	argsLen := len(args)

	if argsLen > 1 {
		for i := 1; i < argsLen; i++ {
			for _, letter := range args[i] {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
