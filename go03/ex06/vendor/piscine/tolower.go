package piscine

func	ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'a' && runes[i] <= 'z') {
			runes[i] += 'A' - 'a';
		}
	}
	ret := string(runes)
	return ret
}