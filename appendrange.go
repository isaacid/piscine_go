package piscine

func AppendRange(min, max int) []int {
	var numSlice []int

	if min > max {
		return []int(nil)
	}

	for i := min; i < max; i++ {
		numSlice = append(numSlice, i)
	}

	return numSlice
}
