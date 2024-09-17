package piscine

func IterativePower(nb int, power int) int {
	result := 1

	if power < 0 {
		result = 0
	} else if power == 0 {
		result = 1
	} else if power == 1 {
		result = nb
	} else {
		for i := power; i > 0; i-- {
			result *= nb
		}
	}
	return result
}
