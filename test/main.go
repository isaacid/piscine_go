// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	argsArr := os.Args
// 	argsArrLen := len(argsArr)

// 	if argsArrLen != 2 || argsArr[1] == "" {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	newStr := ""
// 	newWord := false

// 	for i := 1; i < argsArrLen; i++ {
// 		for _, c := range argsArr[i] {
// 			if c == ' ' && c == '\t' {
// 				newWord = true
// 			} else if c != ' ' && c != '\t' {
// 				newWord = false
// 				newStr += string(c)
// 			}
// 			if newWord {
// 				newStr += " "
// 			}

// 		}

// 	}

// 	for _, c := range newStr {
// 		z01.PrintRune(c)
// 	}

// 	z01.PrintRune('\n')
// }

//

// clean string look into it tonight

//===================

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {

// 	argsArr := os.Args
// 	argsArrLen := len(argsArr)

// 	if argsArrLen != 3 {
// 		return
// 	}

// 	s1 := argsArr[1]
// 	s2 := argsArr[2]

// 	if len(s1) > len(s2) {
// 		return
// 	}

// 	count := 0

// 	for _, c := range s1 {
// 		found := false
// 		for index, s2_c := range s2 {
// 			if c == s2_c {
// 				count++
// 				s2 = s2[index+1:]
// 				found = true
// 				break

// 			}
// 		}

// 		if !found {
// 			return
// 		}

// 	}

// 	if count == len(s1) {
// 		for _, c := range s1 {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')
// 	}

// }

// package piscine

// func ReverseSlice(slice []int) []int {
// 	reversed := make([]int, len(slice))
// 	for i, v := range slice {
// 		reversed[len(slice)-1-i] = v
// 	}
// 	return reversed
// }

// func RevConcatAlternate(slice1, slice2 []int) []int {

// 	s1 := ReverseSlice(slice1)
// 	s2 := ReverseSlice(slice2)

// 	var startSlice, altSlice []int
// 	if len(s1) > len(s2) {
// 		startSlice = s1
// 		altSlice = s2
// 	} else if len(s1) < len(s2) {
// 		startSlice = s2
// 		altSlice = s1
// 	} else {

// 		startSlice = s1
// 		altSlice = s2
// 	}

// 	result := make([]int, 0, len(s1)+len(s2))

// 	for i := 0; i < len(startSlice) || i < len(altSlice); i++ {
// 		if i < len(startSlice) {
// 			result = append(result, startSlice[i])
// 		}
// 		if i < len(altSlice) {
// 			result = append(result, altSlice[i])
// 		}
// 	}

// 	return result
// }

package main

import (
	"fmt"

	"piscine"
)

func main() {
	a1 := []int{-541263, -487372, -304170, -285731, -82786, 323592, 390163, 642540}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(piscine.F, a1)
	result2 := piscine.IsSorted(piscine.F, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
