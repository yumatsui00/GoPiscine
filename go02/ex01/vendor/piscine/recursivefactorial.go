package piscine

func rec(nb int, res int, overcheck int) int {
	if nb == 1{
		return res
	}
	res = res * nb
	if res < overcheck {
		return 0
	}
	overcheck = res
	res = rec(nb - 1, res, overcheck)
	return res
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	res := rec(nb, 1, 0)
	return res
}