package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	res := nb
	for i := 1; i < power; i++ {
		res *= nb
	}
	return res
}