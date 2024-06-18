package piscine

func strlen(a string) int {
	count := 0
	for range a {
		count++
	}
	return count
}

func BasicAtoi(s string) int {
	ret := 0
	l := strlen(s)
	for i := 0; i < l; i++ {
		ret = ret*10 + int(s[i] - '0')
	}
	return int (ret)
}