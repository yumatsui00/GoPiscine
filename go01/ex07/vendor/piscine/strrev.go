package piscine

func strlen(a string) int {
	count := 0
	for range a {
		count++
	}
	return count
}

func StrRev(s string) string {
	l := strlen(s)
	rev := ""
	for i := l - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}