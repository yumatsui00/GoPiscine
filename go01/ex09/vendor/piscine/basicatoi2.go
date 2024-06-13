package piscine

func BasicAtoi2(s string) int {
	ret := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if (s[i] > '9' || s[i] < '0') {
			return 0
		}
		ret = ret*10 + int(s[i] - '0')
	}
	return int (ret)
}