package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *pointStruct) {
	ptr.x = 42
	ptr.y = 21
}

type pointStruct struct {
	x int
	y int
}

func printNbr(nbr int) {
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}

	digits := []rune{}

	for nbr > 0 {
		digit := nbr % 10
		digits = append(digits, rune(digit+'0'))
		nbr /= 10
	}

	if len(digits) == 0 {
		z01.PrintRune('0')
	}

	for i := (len(digits) - 1); i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	point := pointStruct{x: 0, y: 0}

	points := &point

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.y)
	z01.PrintRune('\n')
}
