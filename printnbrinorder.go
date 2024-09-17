package piscine

import (
	"github.com/01-edu/z01"
)

func NumSlice(n int) []int {
	var numbers []int

	if n == 0 {
		return []int{0}
	}

	for n > 0 {
		number := (n % 10)
		numbers = append(numbers, number)
		n /= 10
	}

	return numbers
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func Sort(arr []int) {
	arrLen := len(arr)

	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		nums := NumSlice(n)
		Sort(nums)
		for _, num := range nums {
			z01.PrintRune(rune(num) + '0')
		}
	}
}
