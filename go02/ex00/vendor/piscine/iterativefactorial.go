package piscine

func IterativeFactorial(nb int) int {
	res := 1
	overcheck := 0
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for nb = nb; nb > 0; nb-- {
		res = res * nb;
		if res < overcheck {
			return 0
		}
		overcheck = res
	}
	return res
}