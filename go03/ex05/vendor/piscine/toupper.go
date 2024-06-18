package piscine

import "ft"

func	ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < ft.StrLen(s); i++ {
		if (runes[i] >= 'a' && runes[i] <= 'z') {
			runes[i] += 'A' - 'a';
		}
	}
	ret := string(runes)
	return ret
}