package piscine

func F(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		comparison := f(a[i], a[i+1])

		if comparison > 0 {
			ascending = false
		} else if comparison < 0 {
			descending = false
		}
	}

	return ascending || descending
}
