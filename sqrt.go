package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb; i++ {
		if (nb%i == 0) && (i == nb/i) {
			return i
		}
	}
	return 0
}
