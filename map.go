package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}

	for _, nb := range a {
		result = append(result, f(nb))
	}
	return result
}
