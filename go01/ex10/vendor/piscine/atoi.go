package piscine

func Atoi(s string) int {
	ret := 0
	flag := 1
	l := len(s)
	i := 0
	if s[i] == '+' {
		i ++;
	} else if s[i] == '-' {
		flag = -1;
		i ++;
	}
	for i = i; i < l; i++ {
		if (s[i] > '9' || s[i] < '0') {
			return 0
		}
		ret = ret*10 + int(s[i] - '0')
	}
	return int (flag * ret)
}