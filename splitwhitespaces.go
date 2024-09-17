package piscine

// func SplitWhiteSpaces(s string) {
// 	// var result []string
// 	startPoint := 0

// 	for i, letter := range s {

// 		if letter == ' ' {
// 			newSub := s[startPoint:i]
// 			fmt.Println(newSub)
// 		}

// 		// result = append(result, newSub)

// 		startPoint = i + 1

// 		// if letter == ' ' {
// 		// 	lastSpaceIndex = i
// 		// }

// 		// for j := startPoint; j < lastSpaceIndex; j++ {

// 		// }

// 		// result = append(result, s[startPoint:lastSpaceIndex])

// 		// startPoint = lastSpaceIndex
// 	}

// 	// return []string{"0"}
// }

func SplitWhiteSpaces(s string) []string {
	var result []string
	startPoint := 0

	for i, letter := range s {
		if letter == ' ' {
			newSub := s[startPoint:i]
			if newSub != "" {
				result = append(result, newSub)
			}
			startPoint = i + 1
		}
	}

	if startPoint < len(s) {
		result = append(result, s[startPoint:])
	}

	return result
}

// z01.PrintRune('x')
// z01.PrintRune(' ')
// z01.PrintRune('=')
// z01.PrintRune(' ')
// z01.PrintRune(rune(points.x))
// z01.PrintRune(',')
// z01.PrintRune(' ')
// z01.PrintRune('y')
// z01.PrintRune(' ')
// z01.PrintRune('=')
// z01.PrintRune(' ')
// z01.PrintRune(rune(points.y))
// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
