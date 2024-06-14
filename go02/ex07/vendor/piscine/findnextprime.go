package piscine

func isprime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	}
	for i := 2; i <= nb / i; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for i := nb; true; i++{
		if isprime(i) == true {
			return i
		}
	}
	return (0)
}