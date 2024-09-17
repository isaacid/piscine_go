package piscine

func MakeRange(min, max int) []int {
	if min > max || max == 0 {
		return []int(nil)
	}

	size := max - min // should tell make how big the slice of ints are
	numbsSlice := make([]int, size)

	for i := 0; i < size; i++ {
		numbsSlice[i] = i + min
	}
	return numbsSlice
}
