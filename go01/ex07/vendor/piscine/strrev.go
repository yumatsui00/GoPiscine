package piscine

func StrRev(s string) string {
	l := len(s)
	rev := make([]byte, l)
	for i := 0; i < l; i++ {
		rev[i] = s[l-i-1]
	}
	return string(rev)
}