package piscine

func	ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') {
			runes[i] += 'a' - 'A';
		}
	}
	ret := string(runes)
	return ret
}