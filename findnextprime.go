package piscine

func IsPrime2(nb int) bool {
	if nb <= 1 {
		return false
	}

	if nb <= 3 {
		return true
	}

	if nb%2 == 0 {
		return false
	}

	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}

func FindNextPrime(nb int) int {
	if nb%2 == 0 && nb != 2 {
		nb++
	}

	for i := nb; ; i++ {
		if IsPrime2(i) {
			return i
		}
	}
}
