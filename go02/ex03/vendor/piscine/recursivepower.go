package piscine

func rec(res int, nb int, power int) int {
	if power == 1 {
		return res
	}
	res = rec(res*nb, nb, power - 1)
	return res
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	res := rec(nb, nb, power)
	return res
}