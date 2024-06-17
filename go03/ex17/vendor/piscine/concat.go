package piscine

func	Concat(str1 string, str2 string) string {
	rune1 := []rune(str1)
	rune2 := []rune(str2)
	len1 := len(rune1)
	len2 := len(rune2)
	runes := make([]rune, len1 + len2)
	i := 0
	for i = 0; i < len1; i++ {
		runes[i] = rune1[i]
	}
	for j := 0; j < len2; j++ {
		runes[i + j] = rune2[j]
	}
	return string(runes)
}