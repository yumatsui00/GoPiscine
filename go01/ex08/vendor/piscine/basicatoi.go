package piscine

func BasicAtoi(s string) int {
	ret := 0
	l := len(s)
	for i := 0; i < l; i++ {
		ret = ret*10 + int(s[i] - '0')
	}
	return int (ret)
}