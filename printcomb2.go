package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			// Print the first number
			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0')) // Print the units digit of i

			// Print a space
			z01.PrintRune(' ')

			// Print the second number
			z01.PrintRune(rune(j/10 + '0')) // Print the tens digit of j
			z01.PrintRune(rune(j%10 + '0')) // Print the units digit of j

			// Print a comma and space if this is not the last combination
			if !(i == 98 && j == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n') // Newline at the end for better output formatting
}
